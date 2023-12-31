package inMemoryData

import (
	"fmt"
	"reflect"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Address struct {
	StreetName string `json:"streetName"`
	City       string `json:"city"`
	PostalCode int    `json:"postalCode"`
}
type User struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Password string      `json:"password"`
	Age      int         `json:"age"`
	Misc     interface{} `json:"misc"`
}

type UserInterface interface {
	InsertUser(User) error
	GetAllUsers() []User
	GetSingleUser(int) (User, error)
	UpdateUser(User) error
	DeleteUser(int) error
	SearchUser(map[string]interface{}) []User
}

type Users struct {
	AllUsers []User `json:"users"`
}

func (u *Users) InsertUser(newUser User) error {
	u.AllUsers = append(u.AllUsers, newUser)
	return nil
}

func (u *Users) UpdateUser(newData User) error {
	for i, user := range u.AllUsers {
		if user.ID == newData.ID {
			u.AllUsers[i] = newData
			return nil
		}
	}
	return fmt.Errorf("user Not found it the list with ID :%v", newData.ID)
}

func (u *Users) GetSingleUser(userID int) (User, error) {
	for _, user := range u.AllUsers {
		if user.ID == userID {
			return user, nil
		}
	}
	return User{}, fmt.Errorf("user Not found it the list with ID :%v", userID)
}

func (u *Users) GetAllUsers() []User {
	return u.AllUsers
}

func (u *Users) SearchUser(filter map[string]interface{}) []User {
	var matchedUsers []User
	for _, user := range u.AllUsers {
		if userMatchesFilter(user, filter) {
			matchedUsers = append(matchedUsers, user)
		}
	}
	return matchedUsers
}

func (u *Users) DeleteUser(userID int) error {
	for i, user := range u.AllUsers {
		if user.ID == userID {
			temp := u.AllUsers[i+1:]
			u.AllUsers = u.AllUsers[:i]
			u.AllUsers = append(u.AllUsers, temp...)
			return nil
		}
	}
	return fmt.Errorf("user Not found it the list with ID :%v", userID)
}

func userMatchesFilter(user User, filter map[string]interface{}) bool {
	v := reflect.ValueOf(user)
	caser := cases.Title(language.Und)
	for key, value := range filter {
		// Find the corresponding struct field for each filter key
		field := v.FieldByName(caser.String(key)) // Title case to match struct field name
		fmt.Printf("\nkey : %v , value : %v,  field:%v\n", key, value, field)
		if !field.IsValid() {
			continue // Skip if the field does not exist
		}

		// Compare the field value with the filter value
		if !reflect.DeepEqual(field.Interface(), value) {
			return false // Filter does not match
		}
	}
	return true
}
