package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jairogloz/go-content-manager/pkg/domain"
	"github.com/jairogloz/go-content-manager/pkg/repositories/content_item"
	"github.com/spf13/viper"
)

var config domain.EnvVars

func main() {
	var err error
	config, err = LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %s", err.Error())
	}

	r := gin.Default()

	r.POST("/content", Create)

	log.Fatalln(r.Run(fmt.Sprintf(":%s", config.ServerPort))) // listen and serve on 0.0.0.0:8080 by default
}

func Create(c *gin.Context) {
	// Traducir request
	var contentItemCreateParams domain.ContentItemCreateParams
	if err := c.ShouldBindJSON(&contentItemCreateParams); err != nil {
		log.Println("Failed to bind JSON: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Consumir servicio
	contentItem, err := CreateContentItem(contentItemCreateParams, config)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Traducir respuesta
	c.JSON(200, contentItem)
}

func CreateContentItem(contentItemCreateParams domain.ContentItemCreateParams, config domain.EnvVars) (contentItem *domain.ContentItem, err error) {

	now := time.Now().UTC()
	contentItem = &domain.ContentItem{
		Category:    contentItemCreateParams.Category,
		Description: contentItemCreateParams.Description,
		Title:       contentItemCreateParams.Title,
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}

	err = content_item.InsertContentItem(contentItem, config)
	if err != nil {
		return nil, fmt.Errorf("error inserting content item: %w", err)
	}

	return contentItem, nil
}

func LoadConfig() (config domain.EnvVars, err error) {

	viper.AddConfigPath("../../.")
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	err = viper.ReadInConfig()
	if err != nil {
		return config, fmt.Errorf("error reading config file: %w", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, fmt.Errorf("error unmarshalling config: %w", err)
	}

	return config, nil

}
