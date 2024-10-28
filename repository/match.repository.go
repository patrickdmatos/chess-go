package repository

import (
	"chess_game/models"
)

type MatchRepository interface {
    CreateMatch(match models.Match) error
    GetMatchByID(id string) (models.Match, error)
    UpdateMatchStatus(id string, status string) error
}
