package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	logg "github.com/sirupsen/logrus"
	"github.com/yousefzinsazk78/go_news_app/database"
	"github.com/yousefzinsazk78/go_news_app/database/store"
	"github.com/yousefzinsazk78/go_news_app/handlers"
	"github.com/yousefzinsazk78/go_news_app/middlewares"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	logFile, err := os.OpenFile("./log/general-log.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal("error to open log file in specific path....")
	}

	//setup logrus
	logLevel, err := logg.ParseLevel(os.Getenv("LOGLEVEL"))
	if err != nil {
		logLevel = logg.InfoLevel
	}
	logg.SetLevel(logLevel)
	logg.SetOutput(logFile)
	logg.SetFormatter(&logg.JSONFormatter{})
}

func main() {
	//connect to db
	dbConn, err := sql.Open("postgres", os.Getenv("DATABASEURL"))
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	if err := dbConn.Ping(); err != nil {
		dbConn.Close()
		log.Fatalf("ping error : %s", err)
	}

	var (
		database  = database.NewDatabase(dbConn)
		newsStore = store.NewNewsStore(database)
		handler   = handlers.NewHandler(newsStore)
	)

	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	router.Use(middlewares.LoggingMiddleware())

	//handle simple request
	router.POST("/", handler.HandleNewsPost)
	router.GET("/", handler.HandleNewsGet)

	router.Run(":" + os.Getenv("PORT"))
}
