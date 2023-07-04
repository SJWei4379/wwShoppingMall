package main

import (
	"wwShoppingMall/initialize"
	"wwShoppingMall/router"
)

func init() {
	initialize.Init()
}

func main() {
	r := router.GetEngine()
	if err := r.Run(":8666"); err != nil {
		panic(err)
	}
}
