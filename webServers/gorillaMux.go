package httpServer

import (
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/jaswanth-gorripati/go-examples/inMemoryData"
	log "github.com/sirupsen/logrus"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Infof("Request for Path: %v , Method: %v , Client IP: %v", req.URL.Path, req.Method, req.RemoteAddr)
		next.ServeHTTP(w, req)
	})
}

var gorillaServerAddress = "0.0.0.0:4001"

func StartGorillaMux(wg *sync.WaitGroup) {
	gorillaServer := NewHttpServer{InMemoryData: inMemoryData.Users{AllUsers: []inMemoryData.User{}}}
	router := mux.NewRouter()
	router.Use(Logging)

	apiV1 := router.PathPrefix("/api/v1").Schemes("http").Subrouter()
	users := apiV1.PathPrefix("/user").Subrouter()
	users.HandleFunc("/create", gorillaServer.InsertUser).Methods("POST")
	users.HandleFunc("/update", gorillaServer.UpdateUser).Methods("PATCH")
	users.HandleFunc("/get/{ID}", gorillaServer.GetUserByID).Methods("GET")
	users.HandleFunc("/getAll", gorillaServer.GetAllUsers).Methods("GET")
	users.HandleFunc("/filter", gorillaServer.SearchUsers).Methods("POST")
	users.HandleFunc("/delete", gorillaServer.DeleteUserByID).Methods("DELETE")

	log.Infof("Server started at : %v", gorillaServerAddress)
	err := http.ListenAndServe(gorillaServerAddress, router)
	if err != nil {
		log.Fatalf("Failed to start the server due to %v", err)
	}
	wg.Done()
}
