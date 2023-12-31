package httpServer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/jaswanth-gorripati/go-examples/inMemoryData"
	log "github.com/sirupsen/logrus"
)

var HttpAddress = "0.0.0.0:4000"

func (ns *NewHttpServer) healthCheckResponse(w http.ResponseWriter, req *http.Request) {
	ns.mu.RLock()
	defer ns.mu.RUnlock()
	w.Write([]byte("Server health : ok\n"))
}

func (ns *NewHttpServer) InsertUser(w http.ResponseWriter, req *http.Request) {

	var user inMemoryData.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to read the data from req: %v", err), 400)
		return
	}
	ns.mu.Lock()
	defer ns.mu.Unlock()
	err = ns.InMemoryData.InsertUser(user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to store the data : %v", err), 400)
		return
	}
	w.Write([]byte(fmt.Sprintf("Inserted user with ID :%v", user.ID)))

}

func (ns *NewHttpServer) UpdateUser(w http.ResponseWriter, req *http.Request) {

	var user inMemoryData.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to read the data from req: %v", err), 400)
		return
	}

	ns.mu.Lock()
	defer ns.mu.Unlock()
	err = ns.InMemoryData.UpdateUser(user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to store the data : %v", err), 400)
		return
	}
	w.Write([]byte(fmt.Sprintf("Updated user with ID :%v", user.ID)))
}

func (ns *NewHttpServer) GetUserByID(w http.ResponseWriter, req *http.Request) {

	userIDFromPath := strings.Split(req.URL.Path, "/")[2]
	userID, err := strconv.Atoi(userIDFromPath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid ID :%v", userIDFromPath), 400)
		return
	}
	ns.mu.RLock()
	defer ns.mu.RUnlock()
	user, err := ns.InMemoryData.GetSingleUser(userID)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func (ns *NewHttpServer) GetAllUsers(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	ns.mu.RLock()
	defer ns.mu.RUnlock()
	allUsers := ns.InMemoryData.GetAllUsers()
	json.NewEncoder(w).Encode(allUsers)
}

func (ns *NewHttpServer) SearchUsers(w http.ResponseWriter, req *http.Request) {

	var user map[string]interface{}
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to read the data from req: %v", err), 400)
		return
	}

	ns.mu.RLock()
	defer ns.mu.RUnlock()
	users := ns.InMemoryData.SearchUser(user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to Search the users data : %v", err), 400)
		return
	}
	json.NewEncoder(w).Encode(users)

}

func (ns *NewHttpServer) DeleteUserByID(w http.ResponseWriter, req *http.Request) {

	userIDFromPath := strings.Split(req.URL.Path, "/")[2]
	userID, err := strconv.Atoi(userIDFromPath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid ID :%v", userIDFromPath), 400)
		return
	}

	ns.mu.Lock()
	defer ns.mu.Unlock()
	err = ns.InMemoryData.DeleteUser(userID)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.Write([]byte(fmt.Sprintf("Deleted User With ID : %v\n", userID)))
}

type NewHttpServer struct {
	InMemoryData inMemoryData.Users
	mu           sync.RWMutex
}

func StartHttpServer(wg *sync.WaitGroup) {
	newServer := NewHttpServer{InMemoryData: inMemoryData.Users{AllUsers: []inMemoryData.User{{ID: 1, Name: "One", Password: "12345", Age: 12, Misc: "temp"}}}}
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) { newServer.healthCheckResponse(w, req) })
	http.HandleFunc("/insertUser", newServer.InsertUser)
	http.HandleFunc("/updateUser", newServer.UpdateUser)
	http.HandleFunc("/getUserByID/", newServer.GetUserByID)
	http.HandleFunc("/getAllUsers", newServer.GetAllUsers)
	http.HandleFunc("/searchUsers", newServer.SearchUsers)
	http.HandleFunc("/deleteUserByID/", newServer.DeleteUserByID)

	log.Printf("\nServer started at %v\n", HttpAddress)
	err := http.ListenAndServe(HttpAddress, nil)
	if err != nil {
		log.Fatalf("Failed to start the server : %v\n", err)
	}
	wg.Done()
}
