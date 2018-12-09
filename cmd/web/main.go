package main

import (
	"auto_fertilizer_back/pkg/service"
	"fmt"
)

func main() {
	fmt.Println("Listen start")
	megServer := service.NewShopMessageServer()
	megServer.Run()
}
