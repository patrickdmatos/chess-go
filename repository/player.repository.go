package repository

import (
	"chess_game/models"
	"errors"

	"gorm.io/gorm"
)

type PlayerRepository struct {
	DB *gorm.DB
}

func NewPlayerRepository(db *gorm.DB) *PlayerRepository {
	return &PlayerRepository{DB: db}
}

func (r *PlayerRepository) FindAllPlayers() ([]models.Player, error) {
	var players []models.Player
	result := r.DB.Find(&players)
	return players, result.Error
}

func (r *PlayerRepository) RegisterNewPlayerAccount(player models.Player) (models.Player, error) {
    existsEmail, err := r.EmailExists(player.Email)
    if err != nil {
        return models.Player{}, err
    }
    if existsEmail {
        return models.Player{}, errors.New("email already in use")
    }

    existsUsername, err := r.UsernameExists(player.Username)
    if err != nil {
        return models.Player{}, err
    }
    if existsUsername {
        return models.Player{}, errors.New("username already in use")
    }

    result := r.DB.Create(&player)
    if result.Error != nil {
        return models.Player{}, result.Error
    }
    return player, nil
}

func (r *PlayerRepository) EmailExists(email string) (bool, error) {
    var count int64
    if err := r.DB.Model(&models.Player{}).Where("email = ?", email).Count(&count).Error; err != nil {
        return false, err
    }
    return count > 0, nil
}

func (r *PlayerRepository) UsernameExists(username string) (bool, error) {
    var count int64
    if err := r.DB.Model(&models.Player{}).Where("username = ?", username).Count(&count).Error; err != nil {
        return false, err
    }
    return count > 0, nil
}
