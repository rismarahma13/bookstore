package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rismarahma13/bookstore/controllers"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()

	// Daftarkan rute buku
	router.POST("/books", controllers.CreateBook)
	// Tambahkan rute lainnya jika ada operasi CRUD tambahan: GET, PUT, DELETE

	return router
}
