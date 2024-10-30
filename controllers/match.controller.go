// controllers/match.controller.go
package controllers

import (
	"chess_game/models"
	"chess_game/service"
	"strconv"

	"github.com/gofiber/fiber/v2" // Certifique-se de importar o Fiber
)

type MatchController struct {
	Service *service.MatchService
}

// Método para obter uma partida por ID
func (ctrl *MatchController) GetMatchByID(c *fiber.Ctx) error {
	matchID, err := strconv.ParseUint(c.Params("id"), 10, 32) // Converte o ID da URL para uint
	if err != nil {
		return c.Status(400).SendString("Invalid match ID") // Retorna erro se a conversão falhar
	}

	match, err := ctrl.Service.FindMatchByID(uint(matchID)) // Chama o serviço para obter a partida
	if err != nil {
		return c.Status(404).SendString("Match not found") // Retorna erro se a partida não for encontrada
	}

	return c.Status(200).JSON(match) // Retorna a partida encontrada em formato JSON
}

func (ctrl *MatchController) PostCreateMatch(c *fiber.Ctx) error {
	var match models.Match
	if err := c.BodyParser(&match); err != nil {
        return c.Status(400).SendString("Invalid request body")
    }

	match, err := ctrl.Service.RegisterNewMatch(match)

	if err != nil {
		return c.Status(500).SendString("Error starting match")
	}

	return c.Status(201).JSON(match)

}

func (ctrl *MatchController) PutUpdateStatusMatch(c *fiber.Ctx) error {
	matchID, err := strconv.ParseUint(c.Params("id"), 10, 32) // Converte o ID da URL para uint
	status := c.Params(("status"))
	if err != nil {
		return c.Status(400).SendString("Invalid match ID")
	}

	match, err := ctrl.Service.UpdateStatusMatch(uint(matchID),status)
	if err != nil {
		return c.Status(500).SendString("Match didnt exist!")
	}

	return c.Status(200).JSON(match)
}