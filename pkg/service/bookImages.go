package service

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yeahyeahcore/ChocolateStudio-Api/pkg/models"
)

//GetPhotoBookImage - функция для получения изображения Фотокниги из БД
func (s *Service) GetPhotoBookImage(c *gin.Context) {
	bookImage := &models.Book_image{}

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.String(400, "id required")
		return
	}

	bookImageResult := s.db.First(&bookImage, "id=?", id)
	if bookImageResult.RowsAffected == 0 {
		c.String(404, "bookImage not founded")
	}

	c.JSON(200, bookImage)
}

//DeletePhotoBookImage - функция для удаления изображения Фотокниги из БД
func (s *Service) DeletePhotoBookImage(c *gin.Context) {
	bookImage := &models.Book_image{}

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.String(400, "id required")
		return
	}

	bookImageResult := s.db.Delete(&bookImage, "id=?", id)
	if bookImageResult.RowsAffected == 0 {
		c.String(404, "bookImage not founded")
	}
}

//UpdatePhotoBookImage - функция для обновление изображения Фотокниги из БД
func (s *Service) UpdatePhotoBookImage(c *gin.Context) {
	bookImage := &models.Book_image{}

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.String(400, "id required")
		return
	}

	c.BindJSON(&bookImage)

	bookImageResult := s.db.Where("id=?", id).Updates(&bookImage)
	if bookImageResult.RowsAffected == 0 {
		c.String(404, "bookImage not founded")
	}
}

//InsertPhotoBookImage - функция для добавления изображения Фотокниги в БД
func (s *Service) InsertPhotoBookImage(c *gin.Context) {
	bookImage := &models.Book_image{}

	c.BindJSON(&bookImage)

	bookImageResult := s.db.Create(&bookImage)
	if bookImageResult.RowsAffected == 0 {
		c.String(404, "bookImage not created")
	}
}