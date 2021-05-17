package main

import (
	"github.com/xi102/collection/routers"
	"github.com/xi102/collection/util"
)

func main() {
	util.Config()
	router := routers.InitRouter()
	router.Run(":8080")
}
