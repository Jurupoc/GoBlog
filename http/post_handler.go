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
	response, err := json.Marshal(posts)
	if err != nil {
		 log.Fatal(err.Error())
	}
	
	w.WriteHeader(http.StatusAccepted)
	w.Write(response)
}

func (h Handler) handlerGetPostByID(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	id := strings.TrimPrefix(req.URL.Path, GetPostByIdPath)
	post := h.execution.GetPostByID(id)

	response, err := json.Marshal(post)
	if err != nil {
		 log.Fatal(err.Error())
	}

	w.WriteHeader(http.StatusAccepted)
	w.Write(response)
}

func (h Handler) handlerDeletePostByID(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	id := strings.TrimPrefix(req.URL.Path, DeletePostByIdPath)
	h.execution.DeletePostByID(id)

	w.WriteHeader(http.StatusOK)
}