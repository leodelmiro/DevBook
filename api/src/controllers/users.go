package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, createUserError := io.ReadAll(r.Body)
	if createUserError != nil {
		responses.Error(w, http.StatusUnprocessableEntity, createUserError)
		return
	}

	var user models.User
	if createUserError = json.Unmarshal(requestBody, &user); createUserError != nil {
		responses.Error(w, http.StatusBadRequest, createUserError)
		return
	}

	if createUserError = user.Prepare(); createUserError != nil {
		responses.Error(w, http.StatusBadRequest, createUserError)
		return
	}

	db, createUserError := database.Connect()
	if createUserError != nil {
		responses.Error(w, http.StatusInternalServerError, createUserError)
		return
	}
	defer db.Close()

	repository := repository.NewUserRepository(db)
	user.ID, createUserError = repository.Create(user)
	if createUserError != nil {
		responses.Error(w, http.StatusInternalServerError, createUserError)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, getUserError := database.Connect()
	if getUserError != nil {
		responses.Error(w, http.StatusInternalServerError, getUserError)
		return
	}
	defer db.Close()

	repository := repository.NewUserRepository(db)
	users, getUserError := repository.GetUsersBy(nameOrNick)
	if getUserError != nil {
		responses.Error(w, http.StatusInternalServerError, getUserError)
		return
	}

	responses.JSON(w, http.StatusOK, users)
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
