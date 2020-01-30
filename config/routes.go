package config

import (
	"github.com/gileslloyd/menu-parser/internal/app/base"
	"github.com/gileslloyd/menu-parser/internal/app/controller"
)

var Routes map[string]base.Controller = map[string]base.Controller{
	"menu.parse": controller.NewMenuController(),
}
