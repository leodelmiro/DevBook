package repository

import (
	"api/src/models"
	"database/sql"
)

type Posts struct {
	db *sql.DB
}

func NewPostsRepository(db *sql.DB) *Posts {
	return &Posts{db}
}

func (repository Posts) Create(post models.Post) (uint64, error) {
	statement, err := repository.db.Prepare("insert into posts (title, content, author_id) values (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(post.Title, post.Content, post.AuthorId)
	if err != nil {
		return 0, err
	}

	lastInsertedId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedId), nil
}

func (repository Posts) GetById(postId uint64) (models.Post, error) {
	row, err := repository.db.Query(`
		select p.*, u.nick from 
		posts p inner join users u
		on u.id = p.author_id where p.id = ?
	`, postId)

	if err != nil {
		return models.Post{}, err
	}
	defer row.Close()

	var post models.Post

	if row.Next() {
		if err = row.Scan(
			&post.Id,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return models.Post{}, err
		}
	}

	return post, nil
}

func (repository Posts) Get(userId uint64) ([]models.Post, error) {
	rows, err := repository.db.Query(`
	select distinct p.*, u.nick from posts p 
	inner join users u on u.id = p.author_id 
	inner join followers f on p.author_id = f.user_id
	where u.id = ? or f.follower_id = ?
	order by 1 desc`,
	userId, userId,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post

		if err = rows.Scan(
			&post.Id,
			&post.Title,
			&post.Content,
			&post.AuthorId,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorNick,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}
