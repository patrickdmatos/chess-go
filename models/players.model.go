package models

type Player struct {
    ID       uint   `gorm:"primaryKey;autoIncrement"`
    Username string `gorm:"unique;not null"`
    Email    string `gorm:"unique;not null"`
    Friends  []Friend `gorm:"foreignKey:PlayerID;references:ID"`
    Rating   int    `gorm:"default:0"`
}

type Friend struct {
    ID        uint   `gorm:"primaryKey"`
    PlayerID  uint   // Foreign key
    FriendID  uint   // ID of the friend
}