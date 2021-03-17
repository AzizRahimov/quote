package main

import (
	"fmt"
	"github.com/AzizRahimov/quote/pkg/models"
	app "github.com/AzizRahimov/quote/pkg/server"
	_ "github.com/joho/godotenv/autoload"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"sync/atomic"
	"time"
)

func main() {

	port, found := os.LookupEnv("SERVER_PORT")
	if !found {
		log.Fatal("Error, SERVER_PORT not set")
	}

	router := httprouter.New()
	quoteSvc := models.NewQuotes()
	server := app.NewServer(router, quoteSvc)
	server.Init()
	go worker(time.Minute*5, quoteSvc.DeleteOldQuotes)

	svc := http.Server{
		Handler: server,
		Addr:    port}

	fmt.Println("Server is listening on port", port)
	err := svc.ListenAndServe()
	if err != nil {
		log.Fatal("Error to run Server:", err)
	}

	time.Sleep(time.Minute * 2)
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
