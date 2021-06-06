package server

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/yeahyeahcore/ChocolateStudio-Api/conf"
)

var (
	Host string
	Port string
)

//Start Запускает роутинг
func Start() {

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

	r.Static("/media/photostudio/", "images/photostudio")
	r.Static("/media/photobook/", "images/photobook")

	r.GET("/photography", getPhotography)


	//Server run
	r.Run(fmt.Sprintf("%s:%s", conf.Server.Host, conf.Server.Port)) // listen and serve
}
