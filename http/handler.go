package http

import (
	"goBlog/execution"
	"io"
	"net/http"
)

type Handler struct {
	execution execution.PostExecutorInterface
}

type HandlerInterface interface {
	Server()
}

func NewHandlerInterface(execution execution.PostExecutorInterface) HandlerInterface {
	return &Handler{execution: execution}
}

func (h Handler) Server() {
	http.HandleFunc(DefaultPath, h.handlerDefault)
	http.HandleFunc(SavePostPath, h.handlerSavePost)
	http.HandleFunc(GetPostByIdPath, h.handlerGetPostByID)
	http.HandleFunc(GetAllPostsPath, h.handlerGetAllPosts)

	if err := http.ListenAndServe(":8080", nil); err != http.ErrServerClosed {
		panic(err)
	}
}

func (h Handler) handlerDefault(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello World")
}
