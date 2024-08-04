package controllers

import (
	auth "blog-backend/Auth"
	models "blog-backend/Models"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

/*
* Handle User Login
 */
 func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
    var creds models.Credentials

    // Decode the JSON request body into creds struct
    if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
        var res models.Response
        res.BadResponse(w, http.StatusBadRequest, "Invalid request payload", err)
        
		return
    }

    var user models.User
    foundUser, err := user.FindUserByEmail(server.DB, creds.Email)
    if err != nil {
        var res models.Response
        res.BadResponse(w, http.StatusUnauthorized, "User not found", err)
        
		return
    }

    if err := models.VerifyPassword(foundUser.Password, creds.Password); err != nil {
        var res models.Response
        res.BadResponse(w, http.StatusUnauthorized, "Invalid password", err)
        
		return
    }

    token, err := auth.GenerateToken(foundUser.ID, foundUser.Email, foundUser.Role)
    if err != nil {
        var res models.Response
        res.BadResponse(w, http.StatusInternalServerError, "Failed to generate token", err)
        
		return
    }

    var res models.Response
    res.OkResponse(w, http.StatusOK, "Login successful", map[string]interface{}{"token": token})
}




/*
* Handle User Register
*/
func (server *Server) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	
	// Decode the request body into user model
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		var res models.Response
		res.BadResponse(w, http.StatusBadRequest, "Invalid request payload", err)
		
		return
	}

	log.Printf("Registering user %v", user)

	// Hash the user's password
	if err := user.Prepare(); err != nil {
		var res models.Response
		res.BadResponse(w, http.StatusInternalServerError, "Failed to hash password", err)
		
		return
	}

	// // Check if the user already exists
	// existingUser, err := user.FindUserByEmail(server.DB, user.Email)

	// if err == nil && existingUser.Email != "" {
	// 	var res models.Response
	// 	res.BadResponse(w, http.StatusConflict, "User already exists", nil)
		
	// 	return
	// } else if err != gorm.ErrRecordNotFound {
	// 	// if failed to check if user exists in the DB
	// 	var res models.Response
	// 	res.BadResponse(w, http.StatusInternalServerError, "Failed to check existing user", err)
		
	// 	return
	// }

	

	// Save the new user in the database
	if _, err := user.SaveUser(server.DB); err != nil {
		var res models.Response
		res.BadResponse(w, http.StatusInternalServerError, "Failed to create user", err)
		
		return
	}


	// Generate a JWT token for the newly registered user
	token, err := auth.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		var res models.Response
		res.BadResponse(w, http.StatusInternalServerError, "Failed to generate token", err)
		
		return
	}

	// Send the successful response with the token
	responseData := map[string]interface{}{
		"token": token,
		"user":  user,
	}

	var res models.Response
	res.OkResponse(w, http.StatusCreated, "User registered successfully", responseData)
}



/*
* Handle Single User Delete (only admin)
*/
func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Extract claims from context
	claims, ok := r.Context().Value("claims").(*auth.Claims)
	if !ok {
		var res models.Response
		res.BadResponse(w, http.StatusUnauthorized, "Invalid token", nil)
		return
	}

	// Check if the user is an admin
	var isSAdminUser models.User
	foundUser, err := isSAdminUser.FindUserByEmail(server.DB, claims.Email)
	if err != nil {
		var res models.Response
		res.BadResponse(w, http.StatusUnauthorized, "User not found", err)
		
		return
	}

	//TODO: Check why nil value in BadResponse err field cause 
	if !foundUser.IsAdmin() {
		var res models.Response
		res.BadResponse(w, http.StatusForbidden, "Only admin can delete users", errors.New("not admin user"))
		
		return
	}

	// Extract user ID from URL path
	userID := chi.URLParam(r, "userID")

	// Find the user to delete by ID
	var user models.User
	userToDelete, err := user.FindUserByID(server.DB, userID)
	if err != nil {
		var res models.Response
		res.BadResponse(w, http.StatusNotFound, "User not found", err)
		
		return
	}

	// Delete the user
	if err := user.DeleteUserByID(server.DB, userToDelete.ID); err != nil {
		var res models.Response
		res.BadResponse(w, http.StatusInternalServerError, "Failed to delete user", err)
		
		return
	}

	var res models.Response
	res.OkResponse(w, http.StatusOK, "User deleted successfully", nil)
}





/*
* Handle Get Single User
*/
func (server *Server) GetUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	
	// Use the helper function to find the user by ID
	user := models.User{}
	foundUser, err := user.FindUserByID(server.DB, userID)
	if err != nil {
		var res models.Response
		res.BadResponse(w, http.StatusNotFound, "User not found", err)
		return
	}

	var res models.Response
	res.OkResponse(w, http.StatusOK, "User fetched successfully", map[string]interface{}{"user": foundUser})
}



/*
* Handle Get All Users (only admin)
*/
func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {
	// Extract claims from context
	claims, ok := r.Context().Value("claims").(*auth.Claims)
	if !ok {
		var res models.Response
		res.BadResponse(w, http.StatusUnauthorized, "Invalid token", nil)
		return
	}

	// Check if the user is an admin
	var isSAdminUser models.User
	foundUser, err := isSAdminUser.FindUserByEmail(server.DB, claims.Email)
	if err != nil {
		var res models.Response
		res.BadResponse(w, http.StatusUnauthorized, "User not found", err)
		
		return
	}


	//TODO: Check why nil value in BadResponse err field cause 
	if !foundUser.IsAdmin() {
		var res models.Response
		res.BadResponse(w, http.StatusForbidden, "Only admin can delete users", errors.New("not admin user"))
		
		return
	}

	var users []models.User
	
	if err := server.DB.Find(&users).Error; err != nil {
		var res models.Response
		res.BadResponse(w, http.StatusInternalServerError, "Failed to fetch users", err)
		
		return
	}

	var res models.Response
	res.OkResponse(w, http.StatusOK, "Users fetched successfully", map[string]interface{}{"users": users})
}