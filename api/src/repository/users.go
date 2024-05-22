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

func (repository users) Create(user models.User) (uint64, error) {
	statement, createError := repository.db.Prepare("INSERT INTO users (name, nick, email, password) values (?, ?, ?, ?)")
	if createError != nil {
		return 0, createError
	}
	defer statement.Close()

	result, createError := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if createError != nil {
		return 0, createError
	}

	lastIdInserted, createError := result.LastInsertId()
	if createError != nil {
		return 0, createError
	}

	return uint64(lastIdInserted), nil
}

func (repository users) Get(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) //%nameOrNick%

	rows, getUserByError := repository.db.Query("select id, name, nick, email, createdAt from users where name like ? or nick like ?", nameOrNick, nameOrNick)
	if getUserByError != nil {
		return nil, getUserByError
	}

	defer rows.Close()

	var users = make([]models.User, 0)
	for rows.Next() {
		var user models.User
		if getUserByError = rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); getUserByError != nil {
			return nil, getUserByError
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository users) GetById(id uint64) (models.User, error) {
	row, getUserByIdError := repository.db.Query("select id, name, nick, email, createdAt from users where id = ?", id)

	if getUserByIdError != nil {
		return models.User{}, getUserByIdError
	}

	defer row.Close()

	var user models.User

	if row.Next() {
		if getUserByIdError = row.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); getUserByIdError != nil {
			return models.User{}, getUserByIdError
		}
	}

	return user, nil
}

func (repository users) Update(id uint64, user models.User) error {
	statement, updateError := repository.db.Prepare("update users set name = ?, nick = ?, email = ? where id = ?")

	if updateError != nil {
		return updateError
	}

	defer statement.Close()

	if _, updateError = statement.Exec(user.Name, user.Nick, user.Email, id); updateError != nil {
		return updateError
	}

	return nil
}

func (repository users) Delete(id uint64) error {
	statement, deleteError := repository.db.Prepare("delete from users where id = ?")
	if deleteError != nil {
		return deleteError
	}
	defer statement.Close()

	if _, deleteError = statement.Exec(id); deleteError != nil {
		return deleteError
	}

	return nil
}

func (repository users) GetByEmail(email string) (models.User, error) {
	row, getByEmailError := repository.db.Query("select id, password from users where email = ?", email)
	if getByEmailError != nil {
		return models.User{}, getByEmailError
	}
	defer row.Close()

	var user models.User

	if row.Next() {
		if getByEmailError = row.Scan(&user.ID, &user.Password); getByEmailError != nil {
			return models.User{}, getByEmailError
		}
	}

	return user, nil
}

func (repository users) Follow(userId, followerId uint64) error {
	statement, followError := repository.db.Prepare(
		"INSERT IGNORE INTO followers (user_id, follower_id) values (?, ?)",
	)
	if followError != nil {
		return followError
	}
	defer statement.Close()

	if _, followError = statement.Exec(userId, followerId); followError != nil {
		return followError
	}

	return nil
}

func (repository users) Unfollow(userId, followerId uint64) error {
	statement, unfollowError := repository.db.Prepare(
		"DELETE FROM followers where user_id = ? and follower_id = ?",
	)
	if unfollowError != nil {
		return unfollowError
	}
	defer statement.Close()

	if _, unfollowError = statement.Exec(userId, followerId); unfollowError != nil {
		return unfollowError
	}

	return nil
}

func (repository users) GetFollowers(userId uint64) ([]models.User, error) {
	rows, getFollowersError := repository.db.Query(`
		select u.id, u.name, u.nick, u.email, u.createdAt
		from users u inner join followers f on u.id = f.follower_id where f.user_id = ?
	`, userId)
	if getFollowersError != nil {
		return nil, getFollowersError
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User

		if getFollowersError = rows.Scan(&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); getFollowersError != nil {
			return nil, getFollowersError
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository users) GetFollowing(userId uint64) ([]models.User, error) {
	rows, getFollowingError := repository.db.Query(`
		select u.id, u.name, u.nick, u.email, u.createdAt
		from users u inner join followers f on u.id = f.user_id where  f.follower_id = ?
	`, userId)
	if getFollowingError != nil {
		return nil, getFollowingError
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User

		if getFollowingError = rows.Scan(&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); getFollowingError != nil {
			return nil, getFollowingError
		}

		users = append(users, user)
	}

	return users, nil
}

func (repository users) GetPassword(userId uint64) (string, error) {
	row, getPasswordError := repository.db.Query("select password from users where id = ?", userId)
	if getPasswordError != nil {
		return "", getPasswordError
	}
	defer row.Close()

	var user models.User

	if row.Next() {
		if getPasswordError = row.Scan(&user.Password); getPasswordError != nil {
			return "", getPasswordError
		}
	}

	return user.Password, nil
}

func (repository users) UpdatePassword(userId uint64, password string) error {
	statement, updatePasswordError := repository.db.Prepare("update users set password = ? where id = ?")
	if updatePasswordError != nil {
		return updatePasswordError
	}
	defer statement.Close()

	if _, updatePasswordError = statement.Exec(password, userId); updatePasswordError != nil {
		return updatePasswordError
	}

	return nil
}

