package service

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yeahyeahcore/ChocolateStudio-Api/pkg/models"
)

//GetPhotoBook - функция для получения Фотокниг с Фотками из БД
func (s *Service) GetPhotoBook(c *gin.Context) {
	photoBook := &models.Photobook{}
	bookImage := &[]models.Book_image{}

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.String(400, "id required")
		return
	}

	photoBookResult := s.db.First(&photoBook, "id=?", id)
	if photoBookResult.RowsAffected == 0 {
		c.String(404, "photoBook not founded")
		return
	}

	photosResult := s.db.Find(&bookImage, "book_id=?", id)
	if photosResult.RowsAffected == 0 {
		c.String(404, "photos not founded")
		return
	}

	c.JSON(200, map[string]interface{}{
		"photo":     bookImage,
		"photoBook": photoBook,
	})
}

//DeletePhotoBook - функция для удаления Фотокниги из БД
func (s *Service) DeletePhotoBook(c *gin.Context) {
	photoBook := &models.Photobook{}

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.String(400, "id required")
		return
	}

	photoBookResult := s.db.Delete(&photoBook, "id=?", id)
	if photoBookResult.RowsAffected == 0 {
		c.String(404, "photoBook not founded")
		return
	}

	s.file.RemoveFilesFromFolder(fmt.Sprintf("%s/photobook_%d", s.file.PhotoBookDir, id))
}

//UpdatePhotoBook - функция для обновления Фотокниги из БД
func (s *Service) UpdatePhotoBook(c *gin.Context) {
	photoBook := &models.Photobook{}

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.String(400, "id required")
		return
	}

	c.BindJSON(&photoBook)

	photoBookResult := s.db.Where("id=?", id).Updates(&photoBook)
	if photoBookResult.RowsAffected == 0 {
		c.String(404, "photoBook not founded")
		return
	}
}

//InsertPhotoBook - функция для добавления Фотокниг в БД
func (s *Service) InsertPhotoBook(c *gin.Context) {
	photoBook := &models.Photobook{}

	c.BindJSON(&photoBook)

	photoBookResult := s.db.Create(&photoBook)
	if photoBookResult.RowsAffected == 0 {
		c.String(404, "photoBook not created")
		return
	}
	
}

//GetAllPhotoBook - функция для получения всех Фотокниг с Фотками из БД
func (s *Service) GetAllPhotoBook(c *gin.Context) {
	photoBooks := &[]models.Photobook{}
	bookImages := &[]models.Book_image{}

	photoBookResult := s.db.Find(&photoBooks)
	if photoBookResult.RowsAffected == 0 {
		c.String(404, "photoBook not founded")
		return
	}

	photosResult := s.db.Find(&bookImages)
	if photosResult.RowsAffected == 0 {
		c.String(404, "photos not founded")
		return
	}

	c.JSON(200, map[string]interface{}{
		"photo":     bookImages,
		"photoBook": photoBooks,
	})
}
