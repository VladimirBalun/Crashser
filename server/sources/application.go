package main

import (
	"github.com/go-martini/martini"
	"server/sources/controllers"
)

func main() {
	m := martini.Classic()
	m.Post("/core_dump", controllers.AddCoreDump)
	m.Put("/core_dump_status/:uuid", controllers.AddCoreDump)
	m.Delete("/core_dump/:uuid", controllers.RemoveCoreDump)
	m.Run()
}