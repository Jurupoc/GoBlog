package database

import (
	"database/sql"
	"goBlog/shared"
	"goBlog/models"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type DBConnection struct {
	dbConn *sql.DB
}

func (db *DBConnection) connectSQLiteDB() {
	dbConn, err := sql.Open("sqlite3", shared.DB_DEAFULT_PATH)
	if err != nil {
		log.Fatal(err.Error())
	}

	db.dbConn = dbConn
}

func (db *DBConnection) setupDB() {
	_, err := db.dbConn.Exec(shared.DB_CREATE_TABLE_POSTS)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (db *DBConnection) CloseDB() {
	db.CloseDB()
}

type DBConnectionInterface interface {
	InitDB()
	CloseDB()

	SavePost(models.Post)
	GetAllPosts() []models.Post
	GetPostByID(int) models.Post
}

func NewDBConnectionInterface() DBConnectionInterface {
	return &DBConnection{}
}

func (db *DBConnection) InitDB() {
	db.connectSQLiteDB()
	db.setupDB()
}
