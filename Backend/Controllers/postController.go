package controllers

import (
	auth "blog-backend/Auth"
	models "blog-backend/Models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

/*
* Handle Get All Posts
 */
 func (server *Server) GetPosts(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post
	
	if err := server.DB.Preload("Author").Find(&posts).Error; err != nil {
		var res models.Response
		res.BadResponse(w, http.StatusInternalServerError, "Failed to fetch posts", err)
		
		return
	}

	// Set the password field to an empty string before sending the response to omit from response
	for i := range posts {
		posts[i].Author.Password = ""
	}

	var res models.Response
	res.OkResponse(w, http.StatusOK, "Posts fetched successfully", map[string]interface{}{"posts": posts})
}


/*
* Handle Get Specific User ID Posts based on the requester's id
* TODO: Condition over response json format with/wihtout preLoad author details
 */
 func (server *Server) GetUserPosts(w http.ResponseWriter, r *http.Request) {
	// Extract claims from context
	claims, ok := r.Context().Value("claims").(*auth.Claims)
	if !ok {
		var res models.Response
		res.BadResponse(w, http.StatusUnauthorized, "Invalid token", nil)
		return
	}

	// Since this end point is protected, their must be a bearer token with user id
	userID := claims.UserID
	var posts []models.Post

	
	if err := server.DB.Preload("Author").Where("author_id = ?", userID).Find(&posts).Error; err != nil {
		var res models.Response
		res.BadResponse(w, http.StatusInternalServerError, "Failed to fetch posts", err)
		return
	}

	// Set the password field to an empty string before sending the response to omit from response
	for i := range posts {
		posts[i].Author.Password = ""
	}

	var res models.Response
	res.OkResponse(w, http.StatusOK, "Posts fetched successfully", map[string]interface{}{"posts": posts})
}


/*
* Handle Get Single Post
*/
func (server *Server) GetPost(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "postID")
	
	id, err := strconv.Atoi(postID)
	if err != nil {
		var res models.Response
		res.BadResponse(w, http.StatusBadRequest, "Invalid post ID", err)
		return
	}

	post := models.Post{}
	if err := server.DB.Preload("Author").First(&post, id).Error; err != nil {
		var res models.Response
		res.BadResponse(w, http.StatusNotFound, "Post not found", err)
		return
	}

	// Set the password field to an empty string before sending the response to omit from response
	post.Author.Password = ""

	var res models.Response
	res.OkResponse(w, http.StatusOK, "Post fetched successfully", map[string]interface{}{"post": post})
}



/*
* Handle Create Single Post
*/
func (server *Server) CreatePost(w http.ResponseWriter, r *http.Request) {
	
	// Extract claims from context
	claims, ok := r.Context().Value("claims").(*auth.Claims)
	if !ok {
		var res models.Response
		res.BadResponse(w, http.StatusUnauthorized, "Invalid token", nil)
		return
	}

	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		var res models.Response
		res.BadResponse(w, http.StatusBadRequest, "Invalid request payload", err)
	
		return
	}

	post.AuthorID = claims.UserID
	if _, err := post.SavePost(server.DB); err != nil {
		var res models.Response
		res.BadResponse(w, http.StatusInternalServerError, "Failed to create post", err)
	
		return
	}

	// Retrieve the post with the associated author details for better response craft
	if err := server.DB.Preload("Author").First(&post, post.ID).Error; err != nil {
		var res models.Response
		res.BadResponse(w, http.StatusInternalServerError, "Failed to retrieve post with author details", err)
		return
	}

	// Set the password field to an empty string before sending the response to omit from response
	post.Author.Password = ""
	

	var res models.Response
	res.OkResponse(w, http.StatusCreated, "Post created successfully", map[string]interface{}{"post": post})
}



/*
* Handle Updating Single Post (only admin or post's author)
*/
func (server *Server) UpdatePost(w http.ResponseWriter, r *http.Request) {
	// Extract claims from context
	claims, ok := r.Context().Value("claims").(*auth.Claims)
	if !ok {
		var res models.Response
		res.BadResponse(w, http.StatusUnauthorized, "Invalid token", nil)
		return
	}

	// Decode the request body into a map to get only relevant fields
	var requestBody map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		var res models.Response
		res.BadResponse(w, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	// Retrieve the post by ID from the request URL
	postID := chi.URLParam(r, "postID")
	id, err := strconv.Atoi(postID)
	if err != nil {
		var res models.Response
		res.BadResponse(w, http.StatusBadRequest, "Invalid post ID", err)
		return
	}

	var post models.Post
	if err := server.DB.First(&post, id).Error; err != nil {
		var res models.Response
		res.BadResponse(w, http.StatusNotFound, "Post not found", err)
		return
	}

	// Check if the post belongs to the authenticated user or an admin user
	var user models.User
	if _, err := user.FindUserByID(server.DB, fmt.Sprintf("%d", claims.UserID)); err != nil {
		var res models.Response
		res.BadResponse(w, http.StatusInternalServerError, "Failed to retrieve user", err)
		return
	}

	//TODO: Check why nil value in BadResponse err field cause 
	// "invalid memory address or nil pointer dereference" when this if consditon is true
	if post.AuthorID != claims.UserID && user.Role != "admin" {
		var res models.Response
		res.BadResponse(w, http.StatusForbidden, "You are not authorized to update this post", errors.New("not owner of post"))
		return
	}

	// Update only fields present in the requestBody
	if title, ok := requestBody["title"].(string); ok {
		post.Title = title
	}
	if body, ok := requestBody["body"].(string); ok {
		post.Body = body
	}
	// if img, ok := requestBody["img"].(string); ok {
	// 	post.img = img
	// }
	

	// Save the updated post
	if err := server.DB.Save(&post).Error; err != nil {
		var res models.Response
		res.BadResponse(w, http.StatusInternalServerError, "Failed to update post", err)
		return
	}

	var res models.Response
	res.OkResponse(w, http.StatusOK, "Post updated successfully", map[string]interface{}{"post": post})
}



/*
* Handle Deleting Single Post (only admin or post's author)
*/
func (server *Server) DeletePost(w http.ResponseWriter, r *http.Request) {
	
	// Extract claims from context
	claims, ok := r.Context().Value("claims").(*auth.Claims)
	if !ok {
		var res models.Response
		res.BadResponse(w, http.StatusUnauthorized, "Invalid token", nil)
		return
	}

	// Extract post ID from URL path
	postID := chi.URLParam(r, "postID")

	

	var post models.Post
	if _, err := post.FindPostByID(server.DB, postID); err != nil {
		var res models.Response
		res.BadResponse(w, http.StatusInternalServerError, "No such post id", err)
	
		return
	}

	// Check if the post belongs to the authenticated user or the user is an admin
	if post.AuthorID != claims.UserID && claims.Role != "admin" {
		var res models.Response
		res.BadResponse(w, http.StatusForbidden, "You are not authorized to delete this post", nil)
	
		return
	}

	if err := post.DeletePost(server.DB); err != nil {
		var res models.Response
		res.BadResponse(w, http.StatusInternalServerError, "Failed to delete post", err)
	
		return
	}

	var res models.Response
	res.OkResponse(w, http.StatusOK, "Post deleted successfully", nil)
}
