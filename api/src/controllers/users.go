package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, createUserError := io.ReadAll(r.Body)
	if createUserError != nil {
		log.Fatal(createUserError)
	}

	var user models.User
	if createUserError = json.Unmarshal(requestBody, &user); createUserError != nil {
		log.Fatal(createUserError)
	}

	db, createUserError := database.Connect()
	if createUserError != nil {
		log.Fatal(createUserError)
	}
	repository := repository.NewUserRepository(db)
	userId, createUserError := repository.Create(user)
	if createUserError != nil {
		log.Fatal(createUserError)
	}

	w.Write([]byte(fmt.Sprintf("Inserted id: %d", userId)))
}
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting Users!"))
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting User!"))
}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating User!"))
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting User!"))
}