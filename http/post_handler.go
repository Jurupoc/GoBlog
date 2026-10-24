package http

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

type savePostRequest struct {
	Author  string `json:"author"`
	Content string `json:"content"`
}

func (h Handler) handlerSavePost(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer req.Body.Close()

	var postRequest savePostRequest
	_ = json.NewDecoder(req.Body).Decode(&postRequest)
	_ = h.execution.SaveNewPost(postRequest.Author, postRequest.Content)

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Success")
}

func (h Handler) handlerGetAllPosts(w http.ResponseWriter, _ *http.Request) {
	posts := h.execution.GetAllPosts()
	for _, post := range posts {
		log.Printf("id: %d | author: %s | content: %s | createdAt: %s", post.Id, post.Author, post.Content, post.CreateAt.String())
	}
	w.WriteHeader(http.StatusAccepted)
}

func (h Handler) handlerGetPostByID(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	id := strings.TrimPrefix(req.URL.Path, "/get-post/")
	post := h.execution.GetPostByID(id)
	
	log.Printf("id: %d | author: %s | content: %s | createdAt: %s", post.Id, post.Author, post.Content, post.CreateAt.String())

	w.WriteHeader(http.StatusAccepted)
}
