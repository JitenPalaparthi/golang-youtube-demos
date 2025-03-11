package main

import (
	"demo/handlers"
	"flag"
	"log"
	"net/http"
	"os"
	"runtime"
)

var (
	PORT string
)

func main() {

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello World!"))
	// 	w.WriteHeader(http.StatusOK)
	// })

	// http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Pong"))
	// 	w.WriteHeader(200)
	// })

	PORT = os.Getenv("PORT")
	if PORT == "" {
		//PORT = "8087"
		flag.StringVar(&PORT, "port", "8087", "--port=8087 | -port=8087 | --port 8087")
		flag.Parse()
	}

	http.HandleFunc("/", handlers.Roothandler)

	http.HandleFunc("/ping", handlers.Ping)

	http.HandleFunc("/health", handlers.Health)

	userHandler := new(handlers.Userhandler)

	http.HandleFunc("/users", userHandler.Create)

	// 0.0.0.0:8087 --> this application is reachable on all network interfaces on the system which is running this
	// all network interfaces means --> wifi, bluethooth, localhost, private, publi ,thrunderbolt, ethernet etc..

	log.Println("Application started and listening on port:", PORT)
	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		runtime.Goexit()
	}

}

func Health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ok"))
	w.WriteHeader(200)
}
