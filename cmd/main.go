package main

import (
	"fmt"
	"log"
	"net/http"

	"ascii-art-web/server"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	mux := http.NewServeMux()
	server.New(mux)
	hostAddress := "127.0.0.1"
	hostPort := "3000"
	log.Printf("Listening at http://%s:%s", hostAddress, hostPort)
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", hostAddress, hostPort), mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}
