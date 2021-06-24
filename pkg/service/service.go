package service

import (
	"github.com/gin-gonic/gin"
	"github.com/yeahyeahcore/ChocolateStudio-Api/pkg/file"
	"github.com/yeahyeahcore/ChocolateStudio-Api/pkg/models"
	"github.com/yeahyeahcore/ChocolateStudio-Api/pkg/store"
	"gorm.io/gorm"
)

type Service struct {
	db   *gorm.DB
	file *file.File
	conf *models.Config
}

func NewService(store *store.Store, file *file.File, config *models.Config) *Service {
	return &Service{
		db:   store.DB,
		file: file,
		conf: config,
	}
}

func (s *Service) Set(r *gin.Engine) {

	r.Static(s.conf.HTTP.PhotographyURL, s.conf.HTTP.PhotographyFolder)
	r.Static(s.conf.HTTP.PhotobookURL, s.conf.HTTP.PhotobookFolder)

	//Photography table
	r.GET("/photography", s.GetPhotography)
	r.GET("/photographies", s.GetAllPhotography)
	r.DELETE("/photography", s.DeletePhotography)
	r.PATCH("/photography", s.UpdatePhotography)
	r.POST("/photography", s.InsertPhotography)

	//PhotoBook table
	r.GET("/photobook", s.GetPhotoBook)
	r.GET("/photobooks", s.GetAllPhotoBook)
	r.DELETE("/photobook", s.DeletePhotoBook)
	r.PATCH("/photobook", s.UpdatePhotoBook)
	r.POST("/photobook", s.InsertPhotoBook)

	//Photo_image table
	r.GET("/photoimage", s.GetPhotographyPhoto)
	r.GET("/photoimages", s.GetAllPhotographyPhoto)
	r.DELETE("/photoimage", s.DeletePhotographyPhoto)
	r.POST("/photoimage", s.InsertPhotographyPhoto)
	r.PATCH("/photoimage", s.UpdatePhotographyPhoto)

	//PhotoBook table
	r.GET("/bookimage", s.GetPhotobookImage)
	r.GET("/bookimages", s.GetAllPhotobookImage)
	r.DELETE("/bookimage", s.DeletePhotobookImage)
	r.POST("/bookimage", s.InsertPhotobookImage)
}
