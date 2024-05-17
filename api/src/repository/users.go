package repository

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

	result, createError := statement.Exec(users.Name, users.Nick, users.Email, users.Password)
	if createError != nil {
		return 0, createError
	}

	lastIdInserted, createError := result.LastInsertId()
	if createError != nil {
		return 0, createError
	}

	return uint64(lastIdInserted), nil
}

func (repository users) GetUsersBy(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) //%nameOrNick%

	rows, getUserByError := repository.db.Query("select id, name, nick, email, createdAt from users where name like ? or nick like ?", nameOrNick, nameOrNick)
	if getUserByError != nil {
		return nil, getUserByError
	}

	defer rows.Close()

	var users = make([]models.User , 0)
	for rows.Next() {
		var user models.User
		if getUserByError = rows.Scan(	&user.ID,&user.Name,&user.Nick,&user.Email,&user.CreatedAt) ; getUserByError != nil {
			return nil, getUserByError
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository users) GetUsersById(id uint64) (models.User, error) {
	row, getUserByIdError := repository.db.Query("select id, name, nick, email, createdAt from users where id = ?", id)

	if getUserByIdError != nil {
		return models.User{}, getUserByIdError
	}

	defer row.Close()

	var user models.User

	if row.Next() {
		if getUserByIdError = row.Scan(&user.ID,&user.Name,&user.Nick,&user.Email,&user.CreatedAt); getUserByIdError != nil {
			return models.User{}, getUserByIdError
		}
	}

	return user, nil
}
