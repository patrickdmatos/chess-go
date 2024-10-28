package database

import (
	"chess_game/models" // Substitua pelo caminho correto para seu pacote de modelos

	"gorm.io/driver/sqlite" // Ou outro driver de banco de dados que você estiver usando
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Database é a estrutura que contém a conexão com o banco de dados.
type Database struct {
    DB *gorm.DB
}

// Connect conecta ao banco de dados e realiza a migração.
func Connect() (*Database, error) {
    db, err := gorm.Open(sqlite.Open("chess_game.db"), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })
    if err != nil {
        return nil, err
    }

    // Realiza a migração
    if err := db.AutoMigrate(&models.Player{}); err != nil {
        return nil, err
    }

    return &Database{DB: db}, nil
}
