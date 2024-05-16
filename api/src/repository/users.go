package repository

import (
	"api/src/models"
	"database/sql"
)

type users struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *users {
	return &users{db}
}

func (repository users) Create(users models.User) (uint64, error) {
	statement, createError := repository.db.Prepare("INSERT INTO users (name, nick, email, password) values (?, ?, ?, ?)")
	if createError != nil {
		return 0, createError
	}
	defer statement.Close()

	result , createError := statement.Exec(users.Name, users.Nick, users.Email, users.Password)
	if createError != nil {
		return 0, createError
	}

	lastIdInserted, createError := result.LastInsertId()
	if createError != nil {
		return 0, createError
	}

	return uint64(lastIdInserted), nil
}