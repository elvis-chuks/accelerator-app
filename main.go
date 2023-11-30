package main

import (
	"flag"
	"fmt"
	"inventory/delivery"
	"inventory/pkg/env"
	"inventory/pkg/logger"
	"inventory/repository"
	"log"
	"os"
)

func main() {
	l, err := logger.Init()

	if err != nil {
		panic(err)
	}

	env.LoadConfig(".")

	port := os.Getenv("PORT")

	if port == "" {
		port = "5001"
	}

	repo := repository.SetupDb(l)

	app := delivery.SetupRouter(delivery.Config{
		ProductRepo: repo.ProductRepo,
		SaleRepo:    nil,
	})

	addr := flag.String("addr", fmt.Sprintf(":%s", port), "http service address")
	flag.Parse()
	log.Fatal(app.Listen(*addr))
}
