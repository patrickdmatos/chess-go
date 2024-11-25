package services

import (
	"github.com/patrickdmatos/chess-go-game/internal/models"
	"gorm.io/gorm"
)

// PlayerService estrutura que vai gerenciar o serviço de jogadores
type PlayerService struct {
	DB *gorm.DB
}

// CreatePlayer cria um novo jogador referenciando o usuário no banco
func (service *PlayerService) CreatePlayer(userID uint) (*models.Player, error) {
	// Cria o novo jogador
	player := &models.Player{
		UserID: userID,
		Elo: 1500,
		Role: "player",
	}

	if err := service.DB.Create(player).Error; err != nil {
		return nil, err
	}

	return player, nil
}
