package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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

var config EnvVars

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

	// =================== Save contentItem to database ===================
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(config.MongoDBURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println("Failed to connect to MongoDB: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal_server_error"})
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	collection := client.Database(config.MongoDBName).Collection(config.MongoDBCollNameContentItems)

	ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	_, err = collection.InsertOne(ctx, contentItem)
	if err != nil {
		log.Println("Failed to insert contentItem to MongoDB: ", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal_server_error"})
	}
	// =================== Save contentItem to database ===================

	c.JSON(200, contentItem)
}

type EnvVars struct {
	MongoDBCollNameContentItems string `mapstructure:"MONGO_DB_COLL_NAME_CONTENT_ITEMS"`
	MongoDBName                 string `mapstructure:"MONGO_DB_NAME"`
	MongoDBURI                  string `mapstructure:"MONGO_DB_URI"`
	ServerPort                  string `mapstructure:"SERVER_PORT"`
}

func LoadConfig() (config EnvVars, err error) {

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
