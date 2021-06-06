package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yeahyeahcore/ChocolateStudio-Api/storage"
)

var photographyDB storage.Photography

func getPhotography(c *gin.Context) {
	ID, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(404, gin.H{"error": "неверный формат"})
		return
	}

	photography := photographyDB.Select(ID)

	c.JSON(200, photography)
}
