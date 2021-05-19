package main

import (
	"github.com/xi102/collection/models"
	"github.com/xi102/collection/routers"
	"github.com/xi102/collection/util"
)

func main() {
	util.InitConfig()
	models.InitDatabase()
	router := routers.InitRouter()
	router.Run(":8080")
}
