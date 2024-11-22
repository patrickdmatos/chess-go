package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/patrickdmatos/chess-go-game/internal/database"
	"github.com/patrickdmatos/chess-go-game/internal/players/handlers"
	"github.com/patrickdmatos/chess-go-game/internal/players/services"
	"github.com/patrickdmatos/user-module-go/middleware"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		os.Exit(1)
	}

	// Connect to the database
	db, err := database.ConnectToDatabase()
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		os.Exit(1)
	}

	// Create the player service with the database connection
	playerService := &services.PlayerService{DB: db}

	// Create the player handler
	playerHandler := &handlers.PlayerHandler{PlayerService: playerService}

	// Initialize Fiber app
	app := fiber.New()

	app.Get("/",middleware.Protect(), func(c *fiber.Ctx) error {
        return c.SendString("Serviço de xadrez online!")
    })

	// Route to create a player
	app.Post("/registerPlayer",middleware.Protect(), playerHandler.CreatePlayer)

	// Get the port from environment variable or default to 3000
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("Error: The PORT environment variable is not set!")
		os.Exit(1)
	}

	// Start the server on the specified port
	fmt.Printf("API running on port: %s\n", port)
	if err := app.Listen("0.0.0.0:" + port); err != nil {
		panic(err)
	}
}
