package controllers

import (
	"chess_game/models"
	"chess_game/service"
	"strconv"

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

func (ctrl *PlayerController) PutUpPlayerRating(c *fiber.Ctx) error {
    // Obter o ID do jogador da URL
    idParam := c.Params("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        return c.Status(400).SendString("Invalid player ID")
    }

    // Cria um player com o ID especificado
    player := models.Player{ID: uint(id)}

    // Atualiza o rating do jogador
    updatedPlayer, err := ctrl.Service.Repo.UpdatePlayerRating(player)
    if err != nil {
        return c.Status(500).SendString("Error updating player rating")
    }

    return c.Status(200).JSON(updatedPlayer)
}

func (ctrl *PlayerController) PostAddFriend(c *fiber.Ctx) error {
    playerID, err := strconv.ParseUint(c.Params("playerID"), 10, 32)
    if err != nil {
        return c.Status(400).SendString("Invalid player ID")
    }

    friendID, err := strconv.ParseUint(c.Params("friendID"), 10, 32)
    if err != nil {
        return c.Status(400).SendString("Invalid friend ID")
    }

    addFriend, err := ctrl.Service.AddFriend(uint(playerID),uint(friendID))
    if err != nil {
        return c.Status(500).SendString("Error adding friend, please try again!")
    }

    return c.Status(200).JSON(addFriend)
}

func (ctrl *PlayerController) GetListFriends(c *fiber.Ctx) error {
    playerID, err := strconv.ParseUint(c.Params("playerID"), 10, 32)
    if err != nil {
        return c.Status(400).SendString("Invalid player ID")
    }

    playesFriends, err := ctrl.Service.ListPlayersFriend(uint(playerID))
    if err != nil {
        return c.Status(500).SendString("Error finding friends")
    }

    return c.Status(200).JSON(playesFriends)
}