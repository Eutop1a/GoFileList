package main

import (
	"OJReconstitution/conf"
	"OJReconstitution/routes"
)

func main() {
	conf.Init()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
