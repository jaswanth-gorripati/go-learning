package inMemoryData

import "testing"

func TestInsertUser(t *testing.T) {
	users := Users{AllUsers: []User{{ID: 1, Name: "Test 1", Password: "1234", Misc: "temp"}}}
	newUser := User{ID: 1, Name: "Test 1", Password: "1234", Misc: "temp"}
	err := users.InsertUser(newUser)
	if err != nil {
		t.Errorf("Insert User returned an error :%v", err)
	}
	if len(users.AllUsers) != 2 {
		t.Errorf("Insert operation inserted invalid number of records, %v", len(users.AllUsers))
	}
}
