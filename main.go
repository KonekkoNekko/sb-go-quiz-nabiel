package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sb-go-quiz-nabiel/controllers"
	"sb-go-quiz-nabiel/database"
	"sb-go-quiz-nabiel/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

func main() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8080"
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	database.DBMigrate(DB)

	router := gin.Default()

	api := router.Group("/api")
	api.Use(middleware.Authenticate(DB))
	{
		api.GET("/categories", controllers.HandleFindCategories)
		api.GET("/categories/:id", controllers.HandleFindCategory)
		api.GET("/categories/:id/books", controllers.HandleFindBooksByCategory)
		api.POST("/categories", controllers.HandleCreateCategory)
		api.PUT("/categories/:id", controllers.HandleUpdateCategory)
		api.DELETE("/categories/:id", controllers.HandleDeleteCategory)

		api.GET("/books", controllers.HandleFindBooks)
		api.GET("/books/:id", controllers.HandleFindBook)
		api.POST("/books", controllers.HandleCreateBook)
		api.PUT("/books/:id", controllers.HandleUpdateBook)
		api.DELETE("/books/:id", controllers.HandleDeleteBook)
	}

	router.Run(fmt.Sprintf(":%s", appPort))
}