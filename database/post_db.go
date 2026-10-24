package database

import (
	"goBlog/models"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var (
	INSERT_POST_QUERY    string = `INSERT INTO posts(author, content, createdAt) VALUES(?, ?, ?)`
	GET_ALL_POSTS_QUERY  string = `SELECT * FROM posts LIMIT 50`
	GET_POST_BY_ID_QUERY string = `SELECT * FROM posts WHERE id = ?`
)

func (db *DBConnection) SavePost(post models.Post) {
	_, err := db.dbConn.Exec(INSERT_POST_QUERY, post.Author, post.Content, post.CreateAt)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (db *DBConnection) GetAllPosts() []models.Post {
	rows, err := db.dbConn.Query(GET_ALL_POSTS_QUERY)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	var result []models.Post
	for rows.Next() {
		var post models.Post
		err = rows.Scan(&post.Id, &post.Author, &post.Content, &post.CreateAt)
		if err != nil {
			log.Fatalf("Scan: %v", err.Error())
		}
		result = append(result, post)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func (db *DBConnection) GetPostByID(id int) models.Post {
	row, err := db.dbConn.Query(GET_POST_BY_ID_QUERY, id)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer row.Close()

	if !row.Next() {
		log.Fatal("Query return more than 1 result for query_by_id")
	}
	var result models.Post
	row.Scan(&result.Id, &result.Author, &result.Content, &result.CreateAt)

	return result
}
