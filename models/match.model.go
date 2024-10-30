package models

import "time"

type Match struct {
    ID          uint      `gorm:"primaryKey;autoIncrement"`
    Player1ID   uint      `gorm:"not null"` // ID do primeiro jogador
    Player2ID   uint      `gorm:"not null"` // ID do segundo jogador
    Status      string    `gorm:"default:'in_progress'"` // Status da partida (ex.: 'in_progress', 'completed')
    WinnerID    *uint     `gorm:"null"` // ID do jogador vencedor (null se ainda não finalizada)
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
