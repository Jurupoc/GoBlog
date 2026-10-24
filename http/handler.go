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
	http.HandleFunc("/", h.handlerDefault)
	http.HandleFunc("/post", h.handlerSavePost)
	http.HandleFunc("/get-post/", h.handlerGetPostByID)
	http.HandleFunc("/get-posts", h.handlerGetAllPosts)

	if err := http.ListenAndServe(":8080", nil); err != http.ErrServerClosed {
		panic(err)
	}
}

func (h Handler) handlerDefault(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "Hello World")
}
