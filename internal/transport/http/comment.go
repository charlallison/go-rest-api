package http

import (
	"encoding/json"
	"github.com/charlallison/go-rest-api/internal/comment"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// GetComment - retrieves a comment by Id
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Error parsing ID", err)
		return
	}

	c, err := h.Service.GetComment(uint(i))
	if err != nil {
		sendErrorResponse(w, "Error retrieving comment by id", err)
		return
	}

	_ = sendOKResponse(w, c)
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

	_ = sendOKResponse(w, Response{Message: "Deleted successfully"})
}

// GetAllComments - retrieves all comments from the comment service
func (h *Handler) GetAllComments(w http.ResponseWriter, _ *http.Request) {
	comments, err := h.Service.GetAllComments()
	if err != nil {
		sendErrorResponse(w, "Error getting comments", err)
		return
	}

	_ = sendOKResponse(w, comments)
}

// PostComment - POSTs a comment to the comment service
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var com comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&com); err != nil {
		panic(err)
	}

	c, err := h.Service.PostComment(com)
	if err != nil {
		sendErrorResponse(w, "Could not post new comment", err)
		return
	}

	_ = sendOKResponse(w, c)
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

	var com comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&com); err != nil {
		panic(err)
	}

	com, err = h.Service.UpdateComment(uint(i), com)

	if err != nil {
		sendErrorResponse(w, "Error updating comment", err)
		return
	}

	if err = sendOKResponse(w, com); err != nil {
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
