package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rismarahma13/bookstore/config"
	"github.com/rismarahma13/bookstore/models"
	"github.com/rismarahma13/bookstore/utils"
)

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := config.DB.Create(&book).Error; err != nil {
		utils.InternalServerErrorResponse(c, err)
		return
	}

	utils.SuccessResponse(c, http.StatusCreated, book)
}

// Tambahkan implementasi fungsi CRUD lainnya jika diperlukan (GetBooks, GetBookByID, UpdateBook, DeleteBook)
