package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

/* -------------------------------- Helpers ----------------------------------*/

// Helper function to check if the body is an empty JSON object
func isEmptyJSON(body []byte) bool {
	var obj map[string]interface{}
	if err := json.Unmarshal(body, &obj); err != nil {
		return false
	}
	return len(obj) == 0
}


// Middleware to validate request body
func ValidateRequestBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet && r.Method != http.MethodDelete {
			// Read the body
			body, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Failed to read request body", http.StatusInternalServerError)
				
				return
			}

			// Reset the body so the next handler can read it
			r.Body = io.NopCloser(bytes.NewBuffer(body))

			// Check if the body is empty or contains only an empty JSON object
			if len(body) == 0 || isEmptyJSON(body) {
				http.Error(w, "Request body is required and cannot be empty", http.StatusBadRequest)
				
				return
			}
		}
		
		next.ServeHTTP(w, r)
	})
}



// Middleware to validate JWT token
func AuthMiddleware(next http.Handler) http.Handler {
	
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := VerifyToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			
			return
		}

		// Store claims in the request context if needed
		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
