package main

import (
	"demo/db"
	"demo/handlers"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	PORT string
	DSN  string
)

func main() {

	PORT = os.Getenv("PORT")
	DSN = os.Getenv("DB_CONN")
	if PORT == "" {
		//PORT = "8087"
		flag.StringVar(&PORT, "port", "8087", "--port=8087 | -port=8087 | --port 8087")
	}

	if DSN == "" {
		flag.StringVar(&DSN, "dsn", `host=localhost user=admin password=admin123 dbname=usersdb port=5432 sslmode=disable TimeZone=Asia/Shanghai`, "--dsn= host=localhost user=admin password=admin123 dbname=usersdb port=5432 sslmode=disable TimeZone=Asia/Shanghai | -dsn=host=localhost user=admin password=admin123 dbname=usersdb port=5432 sslmode=disable TimeZone=Asia/Shanghai | --dsn host=localhost user=admin password=admin123 dbname=usersdb port=5432 sslmode=disable TimeZone=Asia/Shanghai")
	}

	flag.Parse()

	e := gin.Default()

	e.GET("/", func(ctx *gin.Context) {
		// String
		// JSON
		ctx.String(http.StatusOK, "Hello World!")
	})

	//str := `Hello \n World\n`

	dbConn, err := db.GetConnection(DSN)
	if err != nil {
		log.Fatalln(err.Error())
	} else {
		log.Println("successfully connected to the database", dbConn)
	}

	e.GET("/ping", handlers.Ping)

	userDb := db.NewUserDB(dbConn)
	uHandler := handlers.NewUserHandler(userDb)
	e.POST("/users", uHandler.Create)
	e.PUT("/users", uHandler.Update)
	e.GET("/users/:id", uHandler.GetByID)
	e.DELETE("/users/:id", uHandler.DeleteByID)
	e.GET("/users/all/:limit/:offset", uHandler.GetAllBy)
	//e.GET("/users/{limit}/{offset}", uHandler.GetAllBy) {} to be given as path param in Gorilla framework but in Gin and Echo just give :id etc..

	e.GET("/users", uHandler.GetAll)

	if err := e.Run(":" + PORT); err != nil { // run the service on this port
		log.Fatalln(err.Error())
	} // http.ListenAndServe

}
