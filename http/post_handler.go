package http

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
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
	if req == nil {
		log.Print("Invalid request")
		io.WriteString(w, "Invalid request")
		return
	}
	defer req.Body.Close()

	var postRequest savePostRequest
	_ = json.NewDecoder(req.Body).Decode(&postRequest)

	_, err := h.execution.SaveNewPost(postRequest.Author, postRequest.Content)
	if err != nil {
		http.Error(w, "Internal Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Success")
}
