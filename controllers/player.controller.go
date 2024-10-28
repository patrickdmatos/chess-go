package controllers

import (
	"chess_game/models"
	"chess_game/service"

	"github.com/gofiber/fiber/v2"
)

type PlayerController struct {
	Service *service.PlayerService
}

func (ctrl *PlayerController) GetPlayers(c *fiber.Ctx) error {
	players, err := ctrl.Service.GetAllPlayersRank()
	if err != nil {
		return c.Status(500).SendString("Error fetching players")
	}
	return c.JSON(players)
}

func (ctrl *PlayerController) PostRegisterPlayer(c *fiber.Ctx) error {
    var player models.Player
    if err := c.BodyParser(&player); err != nil {
        return c.Status(400).SendString("Invalid request body")
    }

    player, err := ctrl.Service.RegisterNewPlayerAccount(player)
    if err != nil {
        switch err.Error() {
        case "email already in use":
            return c.Status(409).SendString(err.Error()) // 409 Conflict
        case "username already in use":
            return c.Status(409).SendString(err.Error()) // 409 Conflict
        default:
            return c.Status(500).SendString("Error registering player")
        }
    }

    return c.Status(201).JSON(player) // Retorne o jogador criado com status 201 Created
}
