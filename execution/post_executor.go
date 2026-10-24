package execution

import (
	"errors"
	"goBlog/database"
	"goBlog/models"
	"log"
)

type postExecutor struct {
	db     database.DBConnectionInterface
}

type PostExecutorInterface interface {
	SaveNewPost(string, string) (models.Post, error)
	GetAllPosts() []models.Post
	GetPostByID(string) (models.Post, error)
}

func NewPostExecutorInterface(db database.DBConnectionInterface) PostExecutorInterface {
	return &postExecutor{
		db:     db,
	}
}

func (executor *postExecutor) SaveNewPost(author, content string) (models.Post, error) {
	err := validatePost(author, content)
	if err != nil {
		log.Fatal(err.Error())
	}
	post := models.NewPost(author, content)
	executor.db.SavePost(post)

	return post, nil
}

func (pe *postExecutor) GetAllPosts() []models.Post {
	return []models.Post{}
}

func (pe *postExecutor) GetPostByID(string) (models.Post, error) {
	return models.Post{}, nil
}

func validatePost(author, content string) error {
	if len(content) == 0 {
		return errors.New("Invalid Author")
	}
	if len(content) == 0 {
		return errors.New("Invalid Content")
	}

	return nil
}
