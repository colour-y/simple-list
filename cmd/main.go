package main

import (
	"simplelist/config"
	"simplelist/pkg/util"
	"simplelist/routes"
)

func main() {
	loading()

	r := routes.NewRouter()
	_ = r.Run(config.HttpPort)
}

func loading() {
	config.Init()
	util.Initlog()

}
