package app

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yeahyeahcore/ChocolateStudio-Api/pkg/file"
	"github.com/yeahyeahcore/ChocolateStudio-Api/pkg/models"
	"github.com/yeahyeahcore/ChocolateStudio-Api/pkg/service"
	"github.com/yeahyeahcore/ChocolateStudio-Api/pkg/store"
)

func Run(configpath string) {
	//создаём объедок конфигурации
	config := &models.Config{}
	err := config.Load(configpath)
	if err != nil {
		log.Fatal(err.Error())
	}

	//создаём объедок файлового обработчика
	//и задаём ему папки с изображениями
	file := file.NewFile(config.HTTP.PhotographyFolder, config.HTTP.PhotobookFolder)

	//создаём объедки хранилища и сервиса
	store := store.NewStore(config)
	service := service.NewService(store, file, config)

	//создаём объедок движка Gin и указываем cors'ы
	r := gin.New()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))

	service.Set(r)

	//запуск
	r.Run(fmt.Sprintf("%s:%s", config.HTTP.Host, config.HTTP.Port))
}
