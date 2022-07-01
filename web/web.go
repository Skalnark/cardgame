package web

import (
	"fmt"
	"net/http"
	service "skalnark/cards/service"

	"gorm.io/gorm"
)

type web struct {
	service service.DeckService
}

func RunServer(db *gorm.DB, hostAddress string) {

	web := &web{
		service: service.NewDeckService(db),
	}
	http.HandleFunc("/deck/create", web.CreateDeck)
	http.HandleFunc("/deck/open", web.OpenDeck)
	http.HandleFunc("/deck/draw", web.DrawCards)

	fmt.Println("Starting server on ", hostAddress)
	fmt.Println(http.ListenAndServe(hostAddress, nil))
}
