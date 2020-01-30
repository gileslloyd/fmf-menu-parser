//+build wireinject

package config

import (
	"github.com/gileslloyd/menu-parser/internal/app/controller"
	"github.com/gileslloyd/menu-parser/internal/domain/menu"
	"github.com/gileslloyd/menu-parser/pkg/infrastructure/cloudvision"
	"github.com/gileslloyd/menu-parser/pkg/infrastructure/delivery/rpc"
	"github.com/google/wire"
)

func CreateMenuController() controller.Menu {
	panic(wire.Build(
		cloudvision.NewMLParser,
		menu.NewService,
		controller.NewMenuController,
	))
}

func CreateMessageListener() rpc.MessageListener {
	panic(wire.Build(
		GetRoutes,
		rpc.NewHandler,
		rpc.NewMessageListener,
	))
}
