package service

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yeahyeahcore/ChocolateStudio-Api/pkg/models"
)

//GetPhotoBookImage - функция для получения изображения Фотокниги из БД
func (s *Service) GetPhotobookImage(c *gin.Context) {
	bookImage := &models.Book_image{}

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.String(400, "id required")
		return
	}

	bookImageResult := s.db.First(&bookImage, "id=?", id)
	if bookImageResult.RowsAffected == 0 {
		c.String(404, "bookImage not founded")
		return
	}

	c.JSON(200, bookImage)
}

//DeletePhotoBookImage - функция для удаления изображения Фотокниги из БД
func (s *Service) DeletePhotobookImage(c *gin.Context) {
	bookImage := &models.Book_image{}

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.String(400, "id required")
		return
	}

	bookImageResult := s.db.Delete(&bookImage, "id=?", id)
	if bookImageResult.RowsAffected == 0 {
		c.String(404, "bookImage not founded")
		return
	}
}

//InsertPhotoBookImage - функция для добавления изображения Фотокниги в БД
func (s *Service) InsertPhotobookImage(c *gin.Context) {
	var form models.ImagesForm
	bookImage := &models.Book_image{}

	if err := c.ShouldBind(&form); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	s.file.SaveFilesToFolder(fmt.Sprintf("%s/photobook_%d", s.file.PhotoBookDir, form.ID), form.Images)

	for _, file := range form.Images {
		bookImage = &models.Book_image{
			BookID:   form.ID,
			ImageURL: fmt.Sprintf("%s/photobook_%d/%s", s.conf.HTTP.PhotobookURL, form.ID, file.Filename),
		}
		bookImageResult := s.db.Create(&bookImage)
		if bookImageResult.RowsAffected == 0 {
			c.String(404, "photoimage not created")
			return
		}
	}
}

//UpdatePhotobookPhoto - функция для обновления изображений Фотокниг в БД
func (s *Service) UpdatePhotobookPhoto(c *gin.Context) {
	var form models.ImagesForm
	bookImage := &models.Book_image{}

	if err := c.ShouldBind(&form); err != nil {
		c.String(http.StatusBadRequest, "bad request")
		return
	}

	bookImageDeleteResult := s.db.Delete(&bookImage, "book_id=?", form.ID)
	if bookImageDeleteResult.RowsAffected != 0 {
		s.file.RemoveFilesFromFolder(fmt.Sprintf("%s/photobook_%d", s.file.PhotoBookDir, form.ID))
	}

	for _, file := range form.Images {
		bookImage = &models.Book_image{
			BookID:   form.ID,
			ImageURL: fmt.Sprintf("%s/photobook_%d/%s", s.conf.HTTP.PhotobookURL, form.ID, file.Filename),
		}
		bookResult := s.db.Create(&bookImage)
		if bookResult.RowsAffected == 0 {
			s.db.Delete(&bookImage, "book_id=?", form.ID)
			c.String(404, "photobookImages not changed")
			return
		}
	}

	s.file.SaveFilesToFolder(fmt.Sprintf("%s/photobook_%d", s.file.PhotoBookDir, form.ID), form.Images)
}

//GetAllPhotobookImage - функция для получения всех изображений Фотокниги из БД
func (s *Service) GetAllPhotobookImage(c *gin.Context) {
	bookImages := &[]models.Book_image{}

	bookImageResult := s.db.Find(&bookImages)
	if bookImageResult.RowsAffected == 0 {
		c.String(404, "bookImage not founded")
		return
	}

	c.JSON(200, bookImages)
}
