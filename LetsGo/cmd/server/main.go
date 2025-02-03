package main

import (
	"github.com/SajjadRezaei/proglog/LetsGo/internal/server"
	"log"
)

func main() {
	ser := server.NewHttpServer(":8080")
	log.Fatal(ser.ListenAndServe())
}
