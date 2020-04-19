package main

import (
	"online_shop/conf"
	"online_shop/router"
)

func main() {
	conf.Init()
	r := router.NewRouter()
	r.Run(":9099")
}
