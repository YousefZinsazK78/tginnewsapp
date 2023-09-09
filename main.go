package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/yousefzinsazk78/go_news_app/database"
	"github.com/yousefzinsazk78/go_news_app/database/store"
	"github.com/yousefzinsazk78/go_news_app/handlers"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

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

	router := gin.Default()
	//handle simple request
	router.POST("/", handler.HandleNewsPost)
	router.GET("/", handler.HandleNewsGet)

	router.Run(":" + os.Getenv("PORT"))
}
