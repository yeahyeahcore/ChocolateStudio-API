package service

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yeahyeahcore/ChocolateStudio-Api/pkg/models"
)

//GetPhotography - функция для получения Фотостудий с Фотками из БД
func (s *Service) GetPhotography(c *gin.Context) {
	photography := &models.Photography{}
	photoimage := &[]models.Photo_image{}

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.String(400, "id required")
		return
	}

	photographyResult := s.db.First(&photography, "id=?", id)
	if photographyResult.RowsAffected == 0 {
		c.String(404, "photography not founded")
		return
	}

	photosResult := s.db.Find(&photoimage, "photo_id=?", id)
	if photosResult.RowsAffected == 0 {
		c.String(404, "photos not founded")
		return
	}

	c.JSON(200, map[string]interface{}{
		"photo":       photoimage,
		"photography": photography,
	})
}

//DeletePhotography - функция для удаления Фотостудии из БД
func (s *Service) DeletePhotography(c *gin.Context) {
	photography := &models.Photography{}

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.String(400, "id required")
		return
	}

	photographyResult := s.db.Delete(&photography, "id=?", id)
	if photographyResult.RowsAffected == 0 {
		c.String(404, "photography not founded")
		return
	}

	s.file.RemoveFilesFromFolder(fmt.Sprintf("%s/photography_%d", s.file.PhotographyDir, id))
}

//UpdatePhotography - функция для обновления Фотостудии из БД
func (s *Service) UpdatePhotography(c *gin.Context) {
	photography := &models.Photography{}

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.String(400, "id required")
		return
	}

	c.BindJSON(&photography)

	photographyResult := s.db.Where("id=?", id).Updates(&photography)
	if photographyResult.RowsAffected == 0 {
		c.String(404, "photography not founded")
		return
	}
}

//InsertPhotography - функция для добавления Фотостудии в БД
func (s *Service) InsertPhotography(c *gin.Context) {
	photography := &models.Photography{}

	c.ShouldBindJSON(&photography)

	photographyResult := s.db.Create(&photography)
	if photographyResult.RowsAffected == 0 {
		c.String(404, "photography not created")
		return
	}
	c.JSON(200, map[string]uint{
		"id": photography.ID,
	})
}

//GetAllPhotography - функция для получения всех Фотостудий с Фотками из БД
func (s *Service) GetAllPhotography(c *gin.Context) {
	photographies := &[]models.Photography{}
	photoimages := &[]models.Photo_image{}

	photographyResult := s.db.Find(&photographies)
	if photographyResult.RowsAffected == 0 {
		c.String(404, "photography not founded")
		return
	}

	photosResult := s.db.Find(&photoimages)
	if photosResult.RowsAffected == 0 {
		c.String(404, "photos not founded")
		return
	}

	c.JSON(200, map[string]interface{}{
		"photo":       photoimages,
		"photography": photographies,
	})
}
