package controller

import (
	"encoding/json"
	"github.com/gileslloyd/menu-parser/internal/domain/menu"
	"github.com/gileslloyd/menu-parser/pkg/infrastructure"
)

type Menu struct {
	service menu.Service
}

func NewMenuController(service menu.Service) Menu {
	return Menu{
		service: service,
	}
}

func (m Menu) Execute(message *infrastructure.Message) (string, error) {
	parsedMenu, err := m.service.ParseMenuImage(message.Get("url", "none"))

	if err != nil {
		return "", err
	}

	response, err := json.Marshal(map[string]string{ "data": parsedMenu.Content})

	if err != nil {
		return "", err
	}

	return string(response), nil
}
