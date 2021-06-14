package service

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yeahyeahcore/ChocolateStudio-Api/pkg/models"
)

//GetPhotographyPhoto - функция для получения изображения Фотостудий из БД
func (s *Service) GetPhotographyPhoto(c *gin.Context) {
	photoimage := &models.Photo_image{}

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.String(400, "id required")
		return
	}

	photoResult := s.db.First(&photoimage, "id=?", id)
	if photoResult.RowsAffected == 0 {
		c.String(404, "photoimage not founded")
	}

	c.JSON(200, photoimage)
}

//DeletePhotographyPhoto - функция для получения изображения Фотостудии из БД
func (s *Service) DeletePhotographyPhoto(c *gin.Context) {
	photoimage := &models.Photo_image{}

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.String(400, "id required")
		return
	}

	photoResult := s.db.Delete(&photoimage, "id=?", id)
	if photoResult.RowsAffected == 0 {
		c.String(404, "photoimage not founded")
	}
}

//UpdatePhotographyPhoto - функция для обновление изображения Фотостудии из БД
func (s *Service) UpdatePhotographyPhoto(c *gin.Context) {
	photoimage := &models.Photo_image{}

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.String(400, "id required")
		return
	}

	c.BindJSON(&photoimage)

	photoResult := s.db.Where("id=?", id).Updates(&photoimage)
	if photoResult.RowsAffected == 0 {
		c.String(404, "photoimage not founded")
	}
}

//InsertPhotographyPhoto - функция для добавления изображения Фотостудии в БД
func (s *Service) InsertPhotographyPhoto(c *gin.Context) {
	photoimage := &models.Photo_image{}

	c.BindJSON(&photoimage)

	photoResult := s.db.Create(&photoimage)
	if photoResult.RowsAffected == 0 {
		c.String(404, "photoimage not created")
	}
}