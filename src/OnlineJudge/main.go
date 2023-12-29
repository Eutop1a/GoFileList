package main

import (
	"OnlineJudge/configs"
	"OnlineJudge/routes"
)

func main() {
	configs.Init()

	r := routes.NewRouter()

	_ = r.Run(configs.HttpPort)

}
