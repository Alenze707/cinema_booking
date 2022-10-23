package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	// Load the postgres driver.
	_ "github.com/lib/pq"

	"cinema-booking.ro/be/pkg/infra/repos"
)

func main() {

	// --------------------------------------------------------
	// For now, all the init (incl. hardcoded values) is here.
	// Later, config will be externalized, and the responsibility
	// of such setup will be in an `App` struct.
	// --------------------------------------------------------

	// Database init.

	dsn := "postgres://postgres:postgres123@localhost:5432/cinema_booking?sslmode=disable"
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalln("Failed to connect to the db:", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalln("Failed to ping the db:", err)
	}
	db.DB.SetMaxOpenConns(3)
	db.DB.SetMaxIdleConns(1)

	// UsersRepo init

	userRepo := repos.NewRoomsRepo(db)
	rooms, err := userRepo.GetAll()
	if err != nil {
		fmt.Println("Failed to fetch the rooms from db:", err)
	}
	fmt.Println("The rooms fetched from db are:")
	for _, r := range rooms {
		fmt.Printf("- id: '%s', name: '%s', seats: %d \n", r.ID, r.Name, r.Seats)
	}
	fmt.Println()

}
