package routes

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/paulktham/CVWO-Project/tree/main/BACK1/internal/database"
)


type Server struct {
	queries *database.Queries
	ctx context.Context
	tokenAuth *jwtauth.JWTAuth
}

func NewServer(q *database.Queries, c context.Context, t *jwtauth.JWTAuth) *Server {
	return &Server{queries: q, ctx: c, tokenAuth: t}
}

func (s *Server) GetAPIRoutes() func(r chi.Router) {
	return func(r chi.Router) {		
		// Public routes
		r.Group(func (r chi.Router) {
			r.Get("/threads", s.getAllThreads())
			r.Get("/threads/{id}", s.getThreadById())
			r.Get("/tags/{tag}", s.getThreadsByTag())
			r.Get("/comments/t-{threadId}", s.getCommentsByThread())
			r.Get("/comments/c-{commentId}", s.getCommentById())
			r.Post("/login", s.login())
			r.Post("/user", s.createUser())		
		})

		// Authenticated routes
		r.Group(func(r chi.Router) {
			// Find JWT tokens
			r.Use(jwtauth.Verifier(s.tokenAuth))

			// Handle valid / invalid tokens with provided middleware
			r.Use(jwtauth.Authenticator)
			r.Post("/tags", s.createTagThread())
			r.Post("/threads", s.createThread())
			r.Post("/comments", s.createComment())
		})
	}
}

//Creating a profile
func (s *Server) createUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user database.CreateUserParams
		
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}

		// Check for existing user
		_, err := s.queries.GetUser(s.ctx, user.Username)
		if err == nil { // User exists
			http.Error(w, "Username already exists!", http.StatusBadRequest)
		} else {
			s.queries.CreateUser(s.ctx, user)
			w.Write([]byte("Successful user creation!"))
		}
	}
}