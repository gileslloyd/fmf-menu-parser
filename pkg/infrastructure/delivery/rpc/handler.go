package rpc

import (
	"encoding/json"
	"fmt"
)

type Handler struct {
	routes map[string][2]string
}

func NewHandler() *Handler {
	return &Handler{
		routes: make(map[string][2]string),
	}
}

func (h Handler) Add(route string, controller string, method string) {
	h.routes[route] = [2]string{
		controller,
		method,
	}
}

func (h Handler) Process(message string) {
	var dat map[string]interface{}

	if err := json.Unmarshal([]byte(message), &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
}
