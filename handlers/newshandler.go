package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yousefzinsazk78/go_news_app/models"
	"github.com/yousefzinsazk78/go_news_app/utils"
)

func (h *handlers) HandleNewsPost(c *gin.Context) {
	var newsPost models.Post
	if err := c.BindJSON(&newsPost); err != nil {
		return
	}
	newsPost.CreatedAt = time.Now()
	err := h.store.InsertNews(&newsPost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}
	c.JSON(http.StatusCreated, gin.H{
		"result": "created successfully",
	})
}

func (h *handlers) HandleNewsGet(c *gin.Context) {
	utils.GeneralLogger.Println("handle news request get method")
	posts, err := h.store.GetNews()
	if err != nil {
		utils.ErrorLogger.Println("error in get news records in database! - " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	utils.GeneralLogger.Println("records retrieved from database successfully!")
	c.JSON(http.StatusOK, gin.H{
		"result": posts,
	})
}
