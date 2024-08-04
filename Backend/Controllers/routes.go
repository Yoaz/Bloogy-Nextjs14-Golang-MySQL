package controllers

import (
	auth "blog-backend/Auth"

	"github.com/go-chi/chi"
)

func (server *Server) InitRoutes() {

    // User Route
	server.Router.Route("/api", func(r chi.Router) {
		r.With(auth.ValidateRequestBody, auth.AuthMiddleware).Route("/users", func(r chi.Router) {
			r.Get("/", server.GetUsers) // No Request Body Validation for GET
			r.Get("/{userID}", server.GetUser) // No Request Body Validation for GET
			r.Delete("/{userID}", server.DeleteUser)
		})

        // Posts Route
		r.With(auth.ValidateRequestBody, auth.AuthMiddleware).Route("/posts", func(r chi.Router) {
			r.Post("/", server.CreatePost)
			r.Get("/", server.GetPosts) // No Request Body Validation for GET
			r.Get("/{postID}", server.GetPost) // No Request Body Validation for GET
			r.Get("/user", server.GetUserPosts) // No request body validation for GET
			r.Put("/{postID}", server.UpdatePost)
			r.Delete("/{postID}", server.DeletePost)
		})

        // Authentication routes
        r.With(auth.ValidateRequestBody).Route("/auth", func(r chi.Router) {
			r.Post("/register", server.Register)
			r.Post("/login", server.Login)
		})
    })
}

