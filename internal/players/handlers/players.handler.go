package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/patrickdmatos/chess-go-game/external/service"
	"github.com/patrickdmatos/chess-go-game/internal/players/services"
)

// PlayerHandler estrutura para o handler de jogador
type PlayerHandler struct {
	PlayerService *services.PlayerService
}

// CreatePlayer cria um jogador no banco de dados
func (handler *PlayerHandler) CreatePlayer(c *fiber.Ctx) error {
	claims, err := service.ValidateTokenWithRemoteAPI(c)

	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).SendString("Erro vincular conta!")
	}
	// Recupera o ID do usuário do contexto
	userID, ok := claims["id"].(float64)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).SendString("ID do usuário não encontrado ou inválido")
	}

	fmt.Println("idValue:", userID) // Verifica o valor de user_id no contexto

	// Chama o serviço para criar o jogador
	player, err := handler.PlayerService.CreatePlayer(uint(userID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erro ao criar jogador: " + err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(player)
}
