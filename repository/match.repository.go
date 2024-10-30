// repository/match_repository.go
package repository

import (
	"chess_game/models"
	"errors"

	"gorm.io/gorm"
)

// Definição da interface MatchRepository
type MatchRepository interface {
	FindMatchByID(id uint) (models.Match, error)
	RegisterNewMatch(match models.Match) (models.Match, error)
    UpdateStatusMatch(id uint, status string) (models.Match, error)
}

type matchRepository struct {
	db *gorm.DB
}

// Função para criar uma nova instância de matchRepository
func NewMatchRepository(db *gorm.DB) MatchRepository {
	return &matchRepository{db: db} // Correção: usar matchRepository
}

// Implementação do método FindMatchByID
func (r *matchRepository) FindMatchByID(id uint) (models.Match, error) {
	var match models.Match
	result := r.db.Where("id = ?", id).First(&match) // Utilize First para obter um único registro
	if result.Error != nil { // Verifique se ocorreu um erro
		return match, result.Error // Retorne o erro, se houver
	}
	return match, nil // Retorne o match encontrado e nil para erro
}

func (r *matchRepository) RegisterNewMatch(match models.Match) (models.Match, error) {
    result := r.db.Create(&match)
    if result.Error != nil {
        return match, result.Error
    }

    return match, nil
}

func (r *matchRepository) UpdateStatusMatch(id uint, status string) (models.Match, error) {
    var match models.Match

    // Busca a partida pelo ID
    result := r.db.First(&match, id)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return match, errors.New("Match not found")
        }
        return match, result.Error // Retorna outros erros, se houver
    }

    // Atualiza o status da partida
    match.Status = status
    if err := r.db.Save(&match).Error; err != nil {
        return match, err // Retorna erro caso ocorra ao salvar
    }

    return match, nil // Retorna o match atualizado e nil para erro
}
