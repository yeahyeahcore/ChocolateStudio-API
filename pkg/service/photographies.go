package service

import (
	"github.com/gin-gonic/gin"
	"github.com/yeahyeahcore/ChocolateStudio-Api/pkg/models"
)

func (s *Service) Photograpy(c *gin.Context) {
	photography := &models.Photography{}
	c.Bind(photography)
	if photography.ID == 0 {
		c.String(400, "id required")
		return
	}
	result := s.db.First(&photography, "id=?", photography.ID)
	if result.RowsAffected == 0 {
		c.String(404, "photography not founded")
	}
	photoimage := &models.Photo_image{}
	s.db.First(&photoimage, "photo_id=?", photography.ID)
	c.JSON(200, map[string]interface{}{
		"photo":       photoimage,
		"photography": photography,
	})
}
