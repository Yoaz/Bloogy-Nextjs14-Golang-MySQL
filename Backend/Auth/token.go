package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

/* -------------------------------- Types ----------------------------------*/

var jwtSecret = []byte(os.Getenv("API_SECRET"))

// Claims structure of the JWT claims
type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"` // Add Role field
	jwt.StandardClaims
}

/* -------------------------------- Helpers ----------------------------------*/

// Generate a JWT token for the given UserID, Email, and Role
func GenerateToken(userID uint, email, role string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // 24 hours expiration time
	claims := &Claims{
		UserID: userID,
		Email:  email,
		Role:   role, // Set Role in the claims
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	return token.SignedString(jwtSecret)
}

// Verifying the JWT token & returns UserID, Email, and Role if valid
func VerifyToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("invalid token signature")
		}
		
		return nil, err
	}

	if !token.Valid {
		
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
