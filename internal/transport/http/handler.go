package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/charlallison/go-rest-api/internal/comment"
	"github.com/gorilla/mux"
)

// Handler - stores pointer to our comment service
type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

// Response - An object to store API response from our API
type Response struct {
	Message string
	Error   string
}

// NewHandler - returns a point to a handler
func NewHandler(s *comment.Service) *Handler {
	return &Handler{
		Service: s,
	}
}

// SetupRoutes - sets up all the routes for our application
func (h *Handler) SetupRoutes() {
	fmt.Println(`Setting up routes`)

	h.Router = mux.NewRouter()

	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comment", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/comment/{id}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("/api/comment/{id}", h.DeleteComment).Methods("DELETE")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		sendOKResponse(w, Response{Message: "I am Alive"})
	})
}

// GetComment - retrieves a comment by Id
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Error parsing ID", err)
		return
	}

	comment, err := h.Service.GetComment(uint(i))
	if err != nil {
		sendErrorResponse(w, "Error retrieving comment by id", err)
		return
	}

	sendOKResponse(w, comment)
}

// DeleteComment - deletes a comment from the database
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Error parsing id", err)
		return
	}

	if err = h.Service.DeleteComment(uint(i)); err != nil {
		sendErrorResponse(w, "Error deleting comment", err)
	}

	sendOKResponse(w, Response{Message: "Deleted successfully"})
}

// GetAllComments - retrieves all comments from the comment service
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.Service.GetAllComments()
	if err != nil {
		sendErrorResponse(w, "Error getting comments", err)
		return
	}

	sendOKResponse(w, comments)
}

// PostComment - POSTs a comment to the comment service
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		panic(err)
	}

	c, err := h.Service.PostComment(comment)
	if err != nil {
		sendErrorResponse(w, "Could not post new comment", err)
		return
	}

	sendOKResponse(w, c)
}

// UpdateComment - Update a comment
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Error parsing id", err)
		return
	}

	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		panic(err)
	}

	comment, err = h.Service.UpdateComment(uint(i), comment)

	if err != nil {
		sendErrorResponse(w, "Error updating comment", err)
		return
	}

	if err = sendOKResponse(w, comment); err != nil {
		panic(err)
	}
}

func sendOKResponse(w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}

func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{
		Message: message,
		Error:   err.Error(),
	}); err != nil {
		panic(err)
	}
}
