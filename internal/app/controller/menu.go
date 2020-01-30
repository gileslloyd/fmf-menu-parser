package controller

import (
	"fmt"

	"github.com/gileslloyd/menu-parser/internal/domain/menu"
	"github.com/gileslloyd/menu-parser/pkg/infrastructure"
	"github.com/gileslloyd/menu-parser/pkg/infrastructure/cloudvision"
)

type Menu struct {
	service menu.Service
}

func NewMenuController() Menu {
	return Menu{
		service: menu.NewService(cloudvision.MlParser{}),
	}
}

func (m Menu) Execute(message *infrastructure.Message) {
	fmt.Println("Menu Controller Received: " + message.Get("url", "none"))
}
