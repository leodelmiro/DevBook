package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	requestBody, loginError := io.ReadAll(r.Body)
	if loginError != nil {
		responses.Error(w, http.StatusUnprocessableEntity, loginError)
		return
	}

	var user models.User
	if loginError = json.Unmarshal(requestBody, &user); loginError != nil {
		responses.Error(w, http.StatusBadRequest, loginError)
		return
	}

	
	db, loginError := database.Connect()
	if loginError != nil {
		responses.Error(w, http.StatusInternalServerError, loginError)
		return
	}
	defer db.Close()

	
	repository := repository.NewUserRepository(db)
	savedUser, loginError := repository.GetByEmail(user.Email)
	if loginError != nil {
		responses.Error(w, http.StatusInternalServerError, loginError)
		return
	}

	if loginError = security.CheckPassword(savedUser.Password, user.Password); loginError != nil {
		responses.Error(w, http.StatusUnauthorized, loginError)
		return
	}

	token, _ := auth.CreateToken(savedUser.ID)
	fmt.Println(token)
}