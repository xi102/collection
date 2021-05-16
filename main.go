package main

import (
	"github.com/xi102/collection/routers"
)

func main() {
	router := routers.InitRouter()
	router.Run(":8080")
}
