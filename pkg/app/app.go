package app

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yeahyeahcore/ChocolateStudio-Api/pkg/models"
	"github.com/yeahyeahcore/ChocolateStudio-Api/pkg/service"
	"github.com/yeahyeahcore/ChocolateStudio-Api/pkg/store"
)

func Run(configpath string) {
	config := &models.Config{}
	err := config.Load(configpath)
	if err != nil {
		log.Fatal(err.Error())
	}
	store := store.NewStore(config)
	service := service.NewService(config, store)
	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))
	service.Set(r)
	r.Run(fmt.Sprintf("%s:%s", config.HTTP.Host, config.HTTP.Port))
}
