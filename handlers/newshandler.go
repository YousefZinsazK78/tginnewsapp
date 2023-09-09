package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yousefzinsazk78/go_news_app/models"
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
			"message": err,
		})
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "created successfully",
	})
}

func (h *handlers) HandleNewsGet(c *gin.Context) {
	posts, err := h.store.GetNews()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"result": posts,
	})
}
