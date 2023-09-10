package main

import (
	"database/sql"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/yousefzinsazk78/go_news_app/database"
	"github.com/yousefzinsazk78/go_news_app/database/store"
	"github.com/yousefzinsazk78/go_news_app/handlers"
)

func TestHandleNewsGet(t *testing.T) {

	dbConn, err := sql.Open("postgres", "postgres://postgres:13781378@localhost:5432/newsappdb?sslmode=disable")
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
	router.GET("/", handler.HandleNewsGet)

	router.Run(":" + os.Getenv("PORT"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	if w.Result().StatusCode != http.StatusOK {
		t.Errorf("expected : 200 , got : %d", w.Result().StatusCode)
	}
}
