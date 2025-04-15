package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	//"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/packetspy/go-payment-gateway/internal/repository"
	"github.com/packetspy/go-payment-gateway/internal/service"
	"github.com/packetspy/go-payment-gateway/internal/web/server"
)

func main() {
	println("Hello World! Payment Gateway!")

	connString := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		GetConnectionString().DB.Host,
		GetConnectionString().DB.Port,
		GetConnectionString().DB.Username,
		GetConnectionString().DB.Password,
		GetConnectionString().DB.Database,
		GetConnectionString().DB.SSLMode,
	)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close()

	accountRepository := repository.NewAccountRepository(db)
	accountService := service.NewAccountService(accountRepository)

	port := os.Getenv("HTTP_PORT")
	server := server.NewServer(accountService, port)

	log.Printf("Starting server on port %s", port)
	if err := server.Start(); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
