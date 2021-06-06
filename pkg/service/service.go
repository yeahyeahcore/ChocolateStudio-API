package service

import (
	"github.com/gin-gonic/gin"
	"github.com/yeahyeahcore/ChocolateStudio-Api/pkg/models"
	"github.com/yeahyeahcore/ChocolateStudio-Api/pkg/store"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(config *models.Config, store *store.Store) *Service {
	return &Service{
		db: store.DB,
	}
}

func (s *Service) Set(r *gin.Engine) {
	r.Static("/media/photostudio/", "images/photostudio")
	r.Static("/media/photobook/", "images/photobook")
	r.GET("/photography", s.Photograpy)

}
