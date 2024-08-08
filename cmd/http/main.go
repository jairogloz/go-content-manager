package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// ContentItemCreateParams is the struct that represents the request body for creating a content item
type ContentItemCreateParams struct {
	Category    string `json:"category"`
	Description string `json:"description"`
	Title       string `json:"title"`
}

// ContentItem reflects a content item in the system.
type ContentItem struct {
	Category              string     `json:"category"`
	CreatedAt             *time.Time `json:"created_at"`
	Description           string     `json:"description"`
	PublishedAtBlog       *time.Time `json:"published_at_blog"`
	PublishedAtLinkedIn   *time.Time `json:"published_at_linkedin"`
	PublishedAtNewsletter *time.Time `json:"published_at_newsletter"`
	PublishedAtYoutube    *time.Time `json:"published_at_youtube"`
	Title                 string     `json:"title"`
	UpdatedAt             *time.Time `json:"updated_at"`
	UserID                string     `json:"-"`
}

func main() {
	r := gin.Default()

	r.POST("/content", Get)

	r.Run() // listen and serve on 0.0.0.0:8080 by default
}

func Get(c *gin.Context) {
	var contentItemCreateParams ContentItemCreateParams
	if err := c.ShouldBindJSON(&contentItemCreateParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	now := time.Now().UTC()
	contentItem := ContentItem{
		Category:    contentItemCreateParams.Category,
		Description: contentItemCreateParams.Description,
		Title:       contentItemCreateParams.Title,
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}

	// Todo: Save contentItem to database

	c.JSON(200, contentItem)
}
