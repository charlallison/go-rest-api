package http

import (
	"github.com/charlallison/go-rest-api/internal/comment"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
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
	log.Info(`Setting up routes`)

	h.Router = mux.NewRouter()
	h.Router.Use(LoggingMiddleware)

	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comment", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/comment/{id}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("/api/comment/{id}", h.DeleteComment).Methods("DELETE")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		_ = sendOKResponse(w, Response{Message: "I am Alive"})
	})
}

// LoggingMiddleware - adds middleware around endpoints
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"Method": r.Method,
			"Path":   r.URL.Path,
		}).Info("handled request")
		next.ServeHTTP(w, r)
	})
}
