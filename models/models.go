package models

import (
	"encoding/json"
	"fmt"
)

// Database struct
type Database struct {
	PostList []Post
}

// Init adds some information in database
func (db *Database) Init() {
	for i := 0; i < 5; i++ {
		db.PostList = append(
			db.PostList,
			Post{
				i,
				fmt.Sprintf("Title of %d post", i),
				fmt.Sprintf("Content of %d post", i)})
	}
}

// Print prints all stuff from database
func (db *Database) Print() {
	for _, post := range db.PostList {
		fmt.Printf("id: %d\ntitle: %s\ncontent: %s\n\n", post.ID, post.Title, post.Content)
	}
}

// GetAll return all stuff from database
func (db *Database) GetAll() ([]byte, error) {
	return json.Marshal(*db)
}

// GetPost get from database post with num
func (db *Database) GetPost(num int) ([]byte, error) {
	for _, post := range db.PostList {
		if post.ID == num {
			return json.Marshal(post)
		}
	}
	return nil, ErrPostNotFound
}

// AddPost ...
func (db *Database) AddPost(title string, content string) {
	lenOfList := len(db.PostList)
	if lenOfList != 0 {
		db.PostList = append(
			db.PostList,
			Post{
				// increment ID of last post in list
				(db.PostList[lenOfList-1]).ID + 1,
				title,
				content})
	}
}
