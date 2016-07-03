// Package authentication provides API to manage user authentication
package authentication

import (
	"encoding/base64"
	"errors"
	"fmt"
	"sync"
)

// type user maintains user list
type user struct {
	username string
	password string
}

var users map[string]user
var mutex = &sync.Mutex{}

func init() {
	users = make(map[string]user)
	users["admin"] = user{
		username: "admin",
		password: base64.StdEncoding.EncodeToString([]byte("admin")),
	}
}

// CheckValidUser determines whether the username and password is valid
func CheckValidUser(username, password string) error {
	if _, ok := users[username]; !ok {
		return fmt.Errorf("username %s is not valid", username)
	}

	//for simplicity sake we use base64 encoding
	passwordBase64 := base64.StdEncoding.EncodeToString([]byte(password))
	if users[username].password != passwordBase64 {
		return fmt.Errorf("password for username %s is not valid", username)
	}

	return nil

}

// AddUser adds new user to the system
func AddUser(username, password string) error {
	mutex.Lock()
	defer mutex.Unlock()
	if _, ok := users[username]; ok {
		return fmt.Errorf("username %s already exists", username)
	}
	if password == "" {
		return errors.New("password must be provided")
	}
	passwordBase64 := base64.StdEncoding.EncodeToString([]byte(password))
	users[username] = user{
		username,
		passwordBase64,
	}
	return nil
}
