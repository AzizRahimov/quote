package main

import (
	"fmt"
	"github.com/AzizRahimov/quote/pkg/models"
	app "github.com/AzizRahimov/quote/pkg/server"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
)

func main() {

	host := "0.0.0.0"
	port := "9999"


	router := httprouter.New()
	quoteSvc := models.NewQuotes()
	server := app.NewServer(router, quoteSvc)
	server.Init()

	svc := http.Server{
		Handler: server,
		Addr: net.JoinHostPort(host, port)}

	fmt.Println("Server is listening port", net.JoinHostPort(host, port))
	err := svc.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
