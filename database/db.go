package database

import (
	"database/sql"
	"goBlog/models"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DBConnection struct {
	dbConn *sql.DB
}

type DBConnectionInterface interface {
	InitDB()
	CloseDB()

	SavePost(models.Post)
	GetAllPosts() []models.Post
	GetPostByID(int) models.Post
	DeletePostByID(id int)
}

func NewDBConnectionInterface() DBConnectionInterface {
	return &DBConnection{}
}

func (db *DBConnection) InitDB() {
	db.connectSQLiteDB()
	db.setupDB()
}

func (db *DBConnection) CloseDB() {
	db.CloseDB()
}

func (db *DBConnection) connectSQLiteDB() {
	dbConn, err := sql.Open("sqlite3", DBDefaultPath)
	if err != nil {
		log.Fatal(err.Error())
	}

	db.dbConn = dbConn
}

func (db *DBConnection) setupDB() {
	_, err := db.dbConn.Exec(DBCreatePostsTable)
	if err != nil {
		log.Fatal(err.Error())
	}
}
