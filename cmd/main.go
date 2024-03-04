package main

import (
	"log"
	"ridhwankiki/bookstore-go/internals/routes"
	"ridhwankiki/bookstore-go/pkg"
)

// Dependency Injection (DI)

func main() {
	// inisialisasi DB
	db, err := pkg.InitMySql()
	if err != nil {
		log.Fatal(err)
		// return
	}
	// inisialisasi Router
	router := routes.InitRouter(db)
	// inisialisasi Server
	server := pkg.InitServer(router)
	// Jalankan Server
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
