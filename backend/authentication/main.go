package main

import (
	"log"
	"mtrain-main/router"
	"mtrain-main/store"
	"mtrain-main/usecases"
	"net/http"
	"os"
)

func main() {
	mongoStore := store.NewMongoDBStore(os.Getenv("MONGO_URI"))

	r := router.NewFiberRouter()
	handler := usecases.NewAccountHandler(mongoStore)

	r.GET("/register", handler.CreateAccount)

	if err := r.Listen(":" + os.Getenv("AUTENTICATION_SERVER_PORT")); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
