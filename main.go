package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	port   int
	folder string
)

func init() {
	flag.IntVar(&port, "port", 3000, "port to listen")
	flag.StringVar(&folder, "folder", ".", "folder to serve")
}

func main() {
	flag.Parse()
	fmt.Printf("Server listening port: %v;\nServe folder: %v ", port, folder)
	err := ServeStatic(port, folder)
	if err != nil {
		log.Fatalln(err)
	}
}

func ServeStatic(port int, folder string) error {
	host := fmt.Sprintf(":%v", port)
	return http.ListenAndServe(host, http.FileServer(http.Dir(folder)))
}
