package services

import (
	"github.com/patrickdmatos/chess-go-game/internal/models"
	"gorm.io/gorm"
)

// PlayerService estrutura que vai gerenciar o serviço de jogadores
type MatchService struct {
	DB *gorm.DB
}


func (service *MatchService) CreateMatch(userID uint) (*models.Player, error) {
	
}