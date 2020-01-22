package main

import (
	"fmt"
	"github.com/gileslloyd/menu-parser/pkg/domain/menu"
	"github.com/gileslloyd/menu-parser/pkg/infrastructure/tensorflow"
)

func main() {
	service := menu.NewService(tensorflow.MlParser{})

	newMenu := service.ParseMenuImage("url")

	fmt.Print("Trader: " + newMenu.TraderID)
	fmt.Print("Content: " + newMenu.Content)
}
