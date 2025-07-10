package config

import (
	"fmt"
	"log"
	"os"

	"github.com/feribeirods/voz-api/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Carrega o .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis do sistema")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Erro ao conectar ao banco de dados: " + err.Error())
	}

	// AutoMigrate com o model de usuário
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		panic("Erro ao migrar banco de dados: " + err.Error())
	}

	DB = db
	log.Println("📦 Banco conectado com sucesso")
}
