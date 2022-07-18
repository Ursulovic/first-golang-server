package database

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Post -
type Post struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UserEmail string    `json:"userEmail"`
	Text      string    `json:"text"`
}

func (c Client) CreatePost(userEmail, text string) (Post, error) {
	db, err := c.readDB()

	if err != nil {
		return Post{}, err
	}

	if _, ok := db.Users[userEmail]; !ok {
		return Post{}, errors.New("user does not exist")
	}

	id := uuid.New().String()

	post := Post{
		ID: id,
		CreatedAt: time.Now().UTC(),
		UserEmail: userEmail,
		Text: text,
	}

	db.Posts[id] = post

	err = c.updateDB(db)

	return post, err
}

func (c Client) GetPosts(userEmail string) ([]Post, error) {

	db, err := c.readDB()

	if err != nil {
		return []Post{}, err
	}

	posts := make([]Post, 0)

	for _, value := range db.Posts {
		if value.UserEmail == userEmail {
			posts = append(posts, value)
		}
	}

	fmt.Println(posts)

	return posts, nil


}

func (c Client) DeletePost(id string) error {

	db, err := c.readDB()

	if err != nil {
		return errors.New("Post not found")
	}

	delete(db.Posts, id)
	
	err = c.updateDB(db)

	return err


}

