package execution

import (
	"errors"
	"goBlog/database"
	"goBlog/models"
	"log"
	"strconv"
)

type postExecutor struct {
	db database.DBConnectionInterface
}

type PostExecutorInterface interface {
	SaveNewPost(string, string) models.Post
	GetAllPosts() []models.Post
	GetPostByID(string) models.Post
}

func NewPostExecutorInterface(db database.DBConnectionInterface) PostExecutorInterface {
	return &postExecutor{
		db: db,
	}
}

func (executor *postExecutor) SaveNewPost(author, content string) models.Post {
	err := validatePost(author, content)
	if err != nil {
		log.Fatal(err.Error())
	}
	post := models.NewPost(author, content)
	executor.db.SavePost(post)

	return post
}

func (pe *postExecutor) GetAllPosts() []models.Post {
	result := pe.db.GetAllPosts()
	return result
}

func (pe *postExecutor) GetPostByID(id string) models.Post {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err.Error())
	}
	result := pe.db.GetPostByID(idInt)
	return result
}

func validatePost(author, content string) error {
	if len(author) == 0 {
		return errors.New("Invalid Author")
	}
	if len(content) == 0 {
		return errors.New("Invalid Content")
	}

	return nil
}
