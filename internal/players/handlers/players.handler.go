package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/patrickdmatos/chess-go-game/internal/players/services"
)

// PlayerHandler estrutura para o handler de jogador
type PlayerHandler struct {
	PlayerService *services.PlayerService
}

// CreatePlayer cria um jogador no banco de dados
func (handler *PlayerHandler) CreatePlayer(c *fiber.Ctx) error {
	// Recupera o ID do usuário do contexto
	userIDValue := c.Locals("user_id")
	fmt.Println("idValue:", userIDValue) // Verifica o valor de user_id no contexto

	if userIDValue == nil {
		return c.Status(fiber.StatusUnauthorized).SendString("ID do usuário não encontrado no contexto")
	}

	// Verifica o tipo de userIDValue e faz a conversão apropriada
	var userID uint
	switch v := userIDValue.(type) {
	case string:
		// Converte string para inteiro
		convertedID, err := strconv.Atoi(v)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString("ID do usuário inválido")
		}
		userID = uint(convertedID)

	case float64:
		// Se for float64 (geralmente JSON pode enviar números assim)
		userID = uint(v)

	default:
		return c.Status(fiber.StatusUnauthorized).SendString("ID do usuário com tipo inesperado")
	}

	fmt.Println("userId:", userID) // Exibe o ID final do usuário para verificação

	// Chama o serviço para criar o jogador
	player, err := handler.PlayerService.CreatePlayer(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Erro ao criar jogador: " + err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(player)
}
