// service/match.service.go
package service

import (
	"chess_game/models"
	"chess_game/repository"
)

type MatchService struct {
	Repo repository.MatchRepository // Alterado para não ser um ponteiro
}

// Função de construção para MatchService
func NewMatchService(repo repository.MatchRepository) *MatchService {
	return &MatchService{Repo: repo} // Agora está correto
}

// Método para obter uma partida por ID
func (s *MatchService) FindMatchByID(id uint) (models.Match, error) {
	return s.Repo.FindMatchByID(id) // Chamada correta do repositório
}

func (s *MatchService) RegisterNewMatch(match models.Match) (models.Match, error) {
    return s.Repo.RegisterNewMatch(match)
}

func (s *MatchService) UpdateStatusMatch(id uint, status string) (models.Match, error) {
    return s.Repo.UpdateStatusMatch(id, status)
}
