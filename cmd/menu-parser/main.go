package main

import "github.com/gileslloyd/menu-parser/internal/app"

func main() {
	//service := menu.NewService(cloudvision.MlParser{})
	//
	//newMenu, err := service.ParseMenuImage("https://media-cdn.tripadvisor.com/media/photo-s/12/f0/f4/1b/menu.jpg")
	//
	//if err == nil {
	//	fmt.Print("Trader: " + newMenu.TraderID)
	//	fmt.Print("Content: " + newMenu.Content)
	//} else {
	//	fmt.Print("Error: " + err.Error())
	//}

	a := app.NewApp()
	a.Run()
}
