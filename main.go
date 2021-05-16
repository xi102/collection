package main

func main() {
	router := InitRouter()
	router.Run(":8080")
}
