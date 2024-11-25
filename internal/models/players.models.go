package models

// Player represents a player in the database
type Player struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	UserID    uint    `gorm:"not null" json:"user_id"`
	GameID    uint    `json:"game_id"`
	Role 	  string  `json:"role"`
	Elo       float64 `gorm:"column:elo_rating" json:"elo"`
}
