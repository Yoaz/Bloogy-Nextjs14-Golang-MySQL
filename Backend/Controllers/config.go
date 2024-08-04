package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	models "blog-backend/Models"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/* -------------------------------- Types ----------------------------------*/

type Server struct {
	DB     *gorm.DB
	Router *chi.Mux
}

/* -------------------------------- Func ----------------------------------*/

// Initialize Database (MySQL) connection using GORM
func (server *Server) InitDB() {
	// Connection string without specifying database
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/"
	dbName := os.Getenv("DB_NAME")

	// Connect to MySQL server to check and create database if necessary
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL server: %v", err)
	}
	defer sqlDB.Close()

	// Create database if it does not exist
	_, err = sqlDB.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s`", dbName))
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}

	// Connection string with database
	dsn = os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	server.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// In case of failure connecting to db
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Check the connection
	sqlDB, err = server.DB.DB()
	if err != nil {
		log.Fatalf("Failed to get database connection: %v", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	// Automigrate models, keep db in sync with structs
	server.DB.AutoMigrate(&models.User{}, &models.Post{})
}

// Initialize Router using Chi
func (server *Server) InitRouter() {
	server.Router = chi.NewRouter() // Create new Chi Router

	// Add CORS middleware
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://frontend:3000"}, 
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	server.Router.Use(c.Handler)

	server.InitRoutes()             // Initialize Routes
}

// Run Server 
func (server *Server) RunServer() {
	addr := os.Getenv("HOST")
	port := os.Getenv("PORT")
	address := fmt.Sprintf("%s:%s", addr, port)
	fmt.Printf("Listening to port %s\n", port)
	log.Fatal(http.ListenAndServe(address, server.Router))
}
