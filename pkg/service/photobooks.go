package service

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yeahyeahcore/ChocolateStudio-Api/pkg/models"
)

//GetPhotoBook - функция для получения Фотокниг с Фотками из БД
func (s *Service) GetPhotoBook(c *gin.Context) {
	photoBook := &models.PhotoBook{}
	bookImage := &models.Book_image{}

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.String(400, "id required")
		return
	}

	photoBookResult := s.db.First(&photoBook, "id=?", id)
	if photoBookResult.RowsAffected == 0 {
		c.String(404, "photoBook not founded")
	}

	photosResult := s.db.First(&bookImage, "photo_id=?", id)
	if photosResult.RowsAffected == 0 {
		c.String(404, "photos not founded")
	}

	c.JSON(200, map[string]interface{}{
		"photo":       photoBook,
		"photoBook": bookImage,
	})
}

//DeletePhotoBook - функция для удаления Фотокниги из БД
func (s *Service) DeletePhotoBook(c *gin.Context) {
	photoBook := &models.PhotoBook{}

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.String(400, "id required")
		return
	}

	photoBookResult := s.db.Delete(&photoBook, "id=?", id)
	if photoBookResult.RowsAffected == 0 {
		c.String(404, "photoBook not founded")
	}
}

//UpdatePhotoBook - функция для обновления Фотокниги из БД
func (s *Service) UpdatePhotoBook(c *gin.Context) {
	photoBook := &models.PhotoBook{}

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.String(400, "id required")
		return
	}

	c.BindJSON(&photoBook)

	photoBookResult := s.db.Where("id=?", id).Updates(&photoBook)
	if photoBookResult.RowsAffected == 0 {
		c.String(404, "photoBook not founded")
	}
}

//InsertPhotoBook - функция для добавления Фотокниг в БД
func (s *Service) InsertPhotoBook(c *gin.Context) {
	photoBook := &models.PhotoBook{}

	c.BindJSON(&photoBook)

	photoBookResult := s.db.Create(&photoBook)
	if photoBookResult.RowsAffected == 0 {
		c.String(404, "photoBook not created")
	}
}