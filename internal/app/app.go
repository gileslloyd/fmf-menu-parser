package app

import (
	"github.com/gileslloyd/menu-parser/config"
	"github.com/gileslloyd/menu-parser/pkg/infrastructure/delivery/rpc"
)

type App struct {
}

func NewApp() App {
	return App{}
}

func (a App) Run() error {
	messageListener := getMessageListener()

	messageListener.Listen()

	return nil
}

func getMessageListener() rpc.MessageListener {
	handler := rpc.NewHandler(config.Routes)

	return rpc.NewMessageListener(handler)
}
