package service

import (
	"fmt"
	"net/http"
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
		return
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
		return
	}
}

//InsertPhotographyPhoto - функция для добавления изображения Фотостудии в БД
func (s *Service) InsertPhotographyPhoto(c *gin.Context) {
	var form models.ImagesForm
	photoimage := &models.Photo_image{}

	if err := c.ShouldBind(&form); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	s.file.SaveFilesToFolder(fmt.Sprintf("%s/photography_%d", s.file.PhotographyDir, form.ID), form.Images)

	for _, file := range form.Images {
		photoimage = &models.Photo_image{
			PhotoID:  form.ID,
			ImageURL: fmt.Sprintf("%s/photography_%d/%s", s.conf.HTTP.PhotographyURL, form.ID, file.Filename),
		}
		photoResult := s.db.Create(&photoimage)
		if photoResult.RowsAffected == 0 {
			c.String(404, "photoimage not created")
			return
		}
	}
}

//GetAllPhotographyPhoto - функция для получения всех изображений Фотостудий из БД
func (s *Service) GetAllPhotographyPhoto(c *gin.Context) {
	photoimages := &models.Photo_image{}

	photoResult := s.db.Find(&photoimages)
	if photoResult.RowsAffected == 0 {
		c.String(404, "photoimage not founded")
		return
	}

	c.JSON(200, photoimages)
}
