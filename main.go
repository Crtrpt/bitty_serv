package main

import (
	"bitty/api"
	"log"
	"net/http"
	"time"

	"bitty/broker"

	"golang.org/x/sync/errgroup"

	"github.com/joho/godotenv"
)

var (
	g errgroup.Group
)


func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}


	

	server01 := &http.Server{
		Addr:         ":9081",
		Handler:      api.Router(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		err := server01.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	g.Go(func() error {
		return broker.Start()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
