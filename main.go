package main

import (
	"chess_game/config"
	"chess_game/controllers"
	"chess_game/database"
	"chess_game/repository"
	"chess_game/routes"
	"chess_game/service"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	 app := fiber.New()

    db, err := database.Connect()
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

	config.LoadEnvVariables() // Load configuration

	// Database connection
	database.Connect()

	// Initialize repositories and services
	playerRepo := repository.NewPlayerRepository(db.DB)
	matchesRepos := repository.NewMatchRepository(db.DB)
	playerService := service.NewPlayerService(playerRepo)
	matchesService := service.NewMatchService(matchesRepos)
	playerController := controllers.PlayerController{Service: playerService}
	matchesController := controllers.MatchController{Service: matchesService}

	// Setup Fiber
	routes.PlayerRoutes(app, &playerController)
	routes.MatchRoutes(app, &matchesController)

	// Start server
	if err := app.Listen(":7777"); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
