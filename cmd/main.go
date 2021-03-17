package main

import (
	"fmt"
	"github.com/AzizRahimov/quote/pkg/models"
	app "github.com/AzizRahimov/quote/pkg/server"
	"github.com/julienschmidt/httprouter"
	"log"
	"net"
	"net/http"
	"sync/atomic"
	"time"
)

func main() {

	host := "0.0.0.0"
	port := "9999"


	router := httprouter.New()
	quoteSvc := models.NewQuotes()
	server := app.NewServer(router, quoteSvc)
	server.Init()
	go worker(time.Minute *5, quoteSvc.DeleteOldQuotes)

	svc := http.Server{
		Handler: server,
		Addr: net.JoinHostPort(host, port)}

	fmt.Println("Server is listening port", net.JoinHostPort(host, port))
	err := svc.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}


func worker(d time.Duration, f func()) {
	var reentrancyFlag int64
	for range time.Tick(d) {
		if atomic.CompareAndSwapInt64(&reentrancyFlag, 0, 1) {
			defer atomic.StoreInt64(&reentrancyFlag, 0)
		} else {
			return
		}
		f()
	}
}
