package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func Connect() {
	godotenv.Load(".env")

	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: os.Getenv("DB"),
	}

	// Get a database handle.
	process(&cfg)
}

func ConnectTesting() {
	godotenv.Load(".env.testing")

	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DBTESTUSER"),
		Passwd: os.Getenv("DBTESTPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: os.Getenv("DBTEST"),
	}

	// Get a database handle.
	process(&cfg)
}

func process(cfg *mysql.Config) {
	// Get a database handle.
	var err error
	DB, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	// fmt.Println("Connected!")
}
