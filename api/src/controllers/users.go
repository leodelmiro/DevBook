package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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

	if createUserError = user.Prepare("create"); createUserError != nil {
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

	db, getUsersError := database.Connect()
	if getUsersError != nil {
		responses.Error(w, http.StatusInternalServerError, getUsersError)
		return
	}
	defer db.Close()

	repository := repository.NewUserRepository(db)
	users, getUsersError := repository.Get(nameOrNick)
	if getUsersError != nil {
		responses.Error(w, http.StatusInternalServerError, getUsersError)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userId, getUserError := strconv.ParseUint(parameters["userId"], 10, 64)
	if getUserError != nil {
		responses.Error(w, http.StatusBadRequest, getUserError)
		return
	}

	db, getUsersError := database.Connect()
	if getUsersError != nil {
		responses.Error(w, http.StatusInternalServerError, getUsersError)
		return
	}
	defer db.Close()

	repository := repository.NewUserRepository(db)
	user, getUserError := repository.GetById(userId)
	if getUserError != nil {
		responses.Error(w, http.StatusInternalServerError, getUserError)
		return
	}

	if user.ID == 0 {
		responses.Error(w, http.StatusNotFound, errors.New("not found"))
		return 
	}

	responses.JSON(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userId, updateUserError := strconv.ParseUint(parameters["userId"], 10, 64)
	if updateUserError != nil {
		responses.Error(w, http.StatusBadRequest, updateUserError)
		return
	}

	requestBody, updateUserError := io.ReadAll(r.Body)
	if updateUserError != nil {
		responses.Error(w, http.StatusUnprocessableEntity, updateUserError)
		return
	}

	var user models.User
	if updateUserError = json.Unmarshal(requestBody, &user); updateUserError != nil {
		responses.Error(w, http.StatusBadRequest, updateUserError)
		return
	}

	if updateUserError = user.Prepare("update"); updateUserError != nil {
		responses.Error(w, http.StatusBadRequest, updateUserError)
		return
	}

	db, getUsersError := database.Connect()
	if getUsersError != nil {
		responses.Error(w, http.StatusInternalServerError, getUsersError)
		return
	}
	defer db.Close()

	repository := repository.NewUserRepository(db)
	if updateUserError := repository.Update(userId, user); updateUserError != nil {
		responses.Error(w, http.StatusInternalServerError, updateUserError)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userId, deleteUserError := strconv.ParseUint(parameters["userId"], 10, 64)
	if deleteUserError != nil {
		responses.Error(w, http.StatusBadRequest, deleteUserError)
		return
	}

	db, getUsersError := database.Connect()
	if getUsersError != nil {
		responses.Error(w, http.StatusInternalServerError, getUsersError)
		return
	}
	defer db.Close()

	repository := repository.NewUserRepository(db)
	if deleteUserError = repository.Delete(userId); deleteUserError != nil {
		responses.Error(w, http.StatusInternalServerError, deleteUserError)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
