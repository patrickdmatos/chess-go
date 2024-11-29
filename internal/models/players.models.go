package models

// Player represents a player in the database
type Player struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	UserID    uint    `gorm:"not null" json:"user_id"`
	GameID    uint    `json:"game_id"`
	Role 	  string  `json:"role"`
	Elo       float64 `gorm:"column:elo_rating" json:"elo"`
}

// TableName define o nome da tabela no banco de dados
func (Player) TableName() string {
	return "players_chess" // Defina o nome da tabela que você deseja
}