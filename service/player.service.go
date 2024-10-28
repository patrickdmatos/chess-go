package service

import (
	"chess_game/models"
	"chess_game/repository"
)

type PlayerService struct {
	Repo *repository.PlayerRepository
}

func NewPlayerService(repo *repository.PlayerRepository) *PlayerService {
	return &PlayerService{Repo: repo}
}

func (s *PlayerService) GetAllPlayersRank() ([]models.Player, error) {
	return s.Repo.FindAllPlayers()
}

// RegisterNewPlayerAccount registers a new player and returns the player
func (s *PlayerService) RegisterNewPlayerAccount(player models.Player) (models.Player, error) {
	return s.Repo.RegisterNewPlayerAccount(player)
}