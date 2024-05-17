package models

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (user *User) Prepare(step string) error {
	if prepareError := user.validate(step); prepareError != nil {
		return prepareError
	}

	user.format()
	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("name is a required field and cannot be empty")
	}	
	
	if user.Nick == "" {
		return errors.New("nick is a required field and cannot be empty")
	}	
	
	if user.Email == "" {
		return errors.New("email is a required field and cannot be empty")
	}

	if validateError := checkmail.ValidateFormat(user.Email); validateError != nil {
		return errors.New("email is invalid format")
	}
	
	if step == "create" && user.Password == "" {
		return errors.New("password is a required field and cannot be empty")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
