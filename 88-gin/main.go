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
	e.Use(LogUrl)
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

	userGroup := e.Group("/users")

	userGroup.Use(LogUrl, IsAutharised, Authenticate)
	userGroup.POST("/", uHandler.Create, PostProcess)
	userGroup.PUT("/", uHandler.Update)
	userGroup.GET("/:id", uHandler.GetByID)
	userGroup.DELETE("/:id", uHandler.DeleteByID)
	userGroup.GET("/all/:limit/:offset", uHandler.GetAllBy)
	//e.GET("/users/{limit}/{offset}", uHandler.GetAllBy) {} to be given as path param in Gorilla framework but in Gin and Echo just give :id etc..
	userGroup.GET("/", uHandler.GetAll)

	if err := e.Run(":" + PORT); err != nil { // run the service on this port
		log.Fatalln(err.Error())
	} // http.ListenAndServe

}

func IsAutharised(ctx *gin.Context) {
	role := ctx.Request.Header.Get("role")
	if role == "admin" {
		ctx.Next() // this moves to the next
	} else {
		ctx.String(http.StatusUnauthorized, "invalid role")
		ctx.Abort()
	}
}

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("token")
	if token == "123456" {
		ctx.Next() // this moves to the next
	} else {
		ctx.String(http.StatusUnauthorized, "invalid token or token is not provided")
		ctx.Abort()
	}
}

func LogUrl(ctx *gin.Context) {
	log.Println((ctx.Request.URL))
	log.Println("clientIP:", ctx.ClientIP())
	log.Println("remoteP:", ctx.RemoteIP())
	ctx.Next()
}

func PostProcess(ctx *gin.Context) {

}
