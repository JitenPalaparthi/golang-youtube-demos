package main

import (
	"demo/handlers"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	PORT string
)

func main() {

	PORT = os.Getenv("PORT")
	if PORT == "" {
		//PORT = "8087"
		flag.StringVar(&PORT, "port", "8087", "--port=8087 | -port=8087 | --port 8087")
		flag.Parse()
	}

	e := gin.Default()

	e.GET("/", func(ctx *gin.Context) {

		// String
		// JSON

		ctx.String(http.StatusOK, "Hello World!")

	})

	e.GET("/ping", handlers.Ping)
	uHandler := handlers.NewUserHandler()
	e.POST("/users", uHandler.Create)

	if err := e.Run(":" + PORT); err != nil { // run the service on this port
		log.Fatalln(err.Error())
	} // http.ListenAndServe

}
