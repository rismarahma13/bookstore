package main

import (
	"github.com/rismarahma13/bookstore/config"
	"github.com/rismarahma13/bookstore/routes"
)

func main() {
	// Koneksi ke database
	config.Connect()

	// Inisialisasi dan jalankan router pada port 8080
	r := routes.InitRoutes()
	if err := r.Run(":8080"); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
