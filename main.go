package main

import (
	"bitty/auth"
	"bitty/docs"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/sync/errgroup"

	"github.com/joho/godotenv"
)

var (
	g errgroup.Group
)

func DocRouter() http.Handler {

	r := gin.New()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Printf(os.Getenv("name"))
	fmt.Printf("=========================================================")
	docs.SwaggerInfo.Host = "127.0.0.1:9081"
	docs.SwaggerInfo.BasePath = "/api/v1"

	g.Go(func() error {
		doc := &http.Server{
			Addr:         ":9080",
			Handler:      DocRouter(),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		}

		err := doc.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		return err
	})

	server01 := &http.Server{
		Addr:         ":9081",
		Handler:      auth.Router(),
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

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
