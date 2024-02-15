package main

import (
	repository "go_final_project/repository"
	router "go_final_project/routers"
)

func main() {
	repository.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
