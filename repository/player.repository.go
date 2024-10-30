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

func (r *PlayerRepository) FindPlayersByRank() ([]models.Player, error) {
	var players []models.Player
	result := r.DB.Order("Rating DESC").Find(&players).Limit(100)
    // Check if any error occurred
    if result.Error != nil {
        return nil, result.Error // Return nil slice and the error
    }

	return players, nil
}

func (r *PlayerRepository) UpdatePlayerRating(player models.Player) (models.Player, error) {
    result := r.DB.Model(&player).Update("Rating", gorm.Expr("Rating + ?", 1))
    return player, result.Error
}

func (r *PlayerRepository) AddFriend(playerID uint, friendID uint) (models.Friend, error) {
    friend := models.Friend{
        PlayerID: playerID,
        FriendID: friendID,
    }
    
    result := r.DB.Create(&friend)
    if result.Error != nil { // Verifica se existe um erro
        return friend, result.Error
    }
    return friend, nil
}

func (r *PlayerRepository) ListPlayersFriend(playerID uint) ([]models.Friend, error) {
    var friends []models.Friend // Use plural for multiple friends
    result := r.DB.Where("player_id = ?", playerID).Find(&friends)

    // Check if any error occurred
    if result.Error != nil {
        return nil, result.Error // Return nil slice and the error
    }

    return friends, nil // Return the slice of friends and no error
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
