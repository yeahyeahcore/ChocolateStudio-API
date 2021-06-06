package main

import (
	"github.com/yeahyeahcore/ChocolateStudio-Api/conf"
	"github.com/yeahyeahcore/ChocolateStudio-Api/server"
	"github.com/yeahyeahcore/ChocolateStudio-Api/storage"
)

func main() {
	conf.Load("config.json")
	storage.Init()
	server.Start()
}
