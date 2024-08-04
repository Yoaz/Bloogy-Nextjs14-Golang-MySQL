package models

import (
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

/* -------------------------------- Types ----------------------------------*/

type User struct {
    gorm.Model
    Email    string `json:"email" gorm:"unique"`
    Password string `json:"password"`
    Role     string `json:"role"` // "user" or "admin"
}


/* -------------------------------- Helpers ----------------------------------*/
// Hash password using built-in Golang package bcrypt
func Hash(pwd string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
}

// Verify password hash against plain password string
func VerifyPassword(hashedPwd string, pwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd))
}

// Before adding user to DB
func (u *User) Prepare() (err error) {
	hashedPassword, err := Hash(u.Password)
	
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	
	return nil
}

// Check if the user is an admin
func (u *User) IsAdmin() bool {
	return u.Role == "admin"
}

/* -------------------------------- DB Actions ----------------------------------*/

// Save user to the db
func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	
	// Log the user data for debugging
    log.Printf("Saving user: %+v\n", u)
	
	err := db.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	
	return u, nil
}



// Find User By Email
func (u *User) FindUserByEmail(db *gorm.DB, email string) (*User, error) {
	
	
	// Log the email for debugging
	log.Printf("Finding user by email: %s\n", email)

	err := db.Where("email = ?", email).First(&u).Error
	if err != nil {
		return &User{}, err
	}
	
	return u, nil
}



// Find User By ID
func (u *User) FindUserByID(db *gorm.DB, id string) (*User, error) {

	// Log the id for debugging
	log.Printf("Finding user by id: %s\n", id)

	err := db.Where("id = ?", id).First(&u).Error
	if err != nil {
		return &User{}, err
	}
	
	return u, nil
}



// Delete User By ID
func (u *User) DeleteUserByID(db *gorm.DB, id uint) error {
	return db.Delete(&User{}, id).Error
}
