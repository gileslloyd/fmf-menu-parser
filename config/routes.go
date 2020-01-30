package config

import (
	"github.com/gileslloyd/menu-parser/internal/app/base"
)

var Routes map[string]base.Controller = map[string]base.Controller{
	"menu.parse": CreateMenuController(),
}
