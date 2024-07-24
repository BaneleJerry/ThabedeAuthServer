package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/BaneleJerry/ThabedeAuthServer/config"
	internalDatabase "github.com/BaneleJerry/ThabedeAuthServer/internal/interfaces/database"
	"github.com/BaneleJerry/ThabedeAuthServer/internal/server"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	database, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	dbQueries := internalDatabase.New(database)

	serverCfg := config.ServerConfig{
		Port:           port,
		IdleTimeout:    time.Minute,
		ReadTimeout:    2 * time.Minute,
		WriteTimeout:   2 * time.Minute,
		MaxHeaderBytes: 1 << 20,
	}

	srv := server.NewServer(dbQueries, serverCfg)
	srv.Start()
}
