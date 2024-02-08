package main

import (
	"log"
	"only-test/router"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	gin.SetMode(os.Getenv("GIN_MODE"))

	r := router.Route()
	r.Run("0.0.0.0:8080")
	// log.Print("Starting")
}
