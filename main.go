package main

import (
	"todo-api/routes"
)

func main() {

	r := routes.TodoRoutes()
	r.Run()
}
