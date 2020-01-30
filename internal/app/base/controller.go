package base

import (
	"github.com/gileslloyd/menu-parser/pkg/infrastructure"
)

type Controller interface {
	Execute(message *infrastructure.Message)
}
