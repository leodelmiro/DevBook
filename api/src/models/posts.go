package models

import (
	"errors"
	"strings"
	"time"
)

type Post struct {
	Id         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorId   uint64    `json:"autorId,omitempty"`
	AuthorNick string    `json:"autorNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
}

func (post *Post) Prepare() error {
	if err := post.validate(); err != nil {
		return err
	}

	if err := post.format(); err != nil {
		return err
	}

	return nil
}

func (post *Post) validate() error {
	if post.Title == "" {
		return errors.New("title is a required field and cannot be empty")
	}

	if post.Content == "" {
		return errors.New("content is a required field and cannot be empty")
	}

	return nil
}

func (post *Post) format() error {
	post.Title = strings.TrimSpace(post.Title)
	post.Content = strings.TrimSpace(post.Content)

	return nil
}