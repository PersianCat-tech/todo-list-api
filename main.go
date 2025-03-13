package main

import (
	"todo_list/conf"
	"todo_list/routes"
)

func main() {
	conf.Init()
	r := routes.NewRouter()
	r.Run(conf.HttpPort)
}