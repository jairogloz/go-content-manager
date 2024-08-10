package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jairogloz/go-content-manager/pkg/domain"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

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
	// Traducir request
	var contentItemCreateParams domain.ContentItemCreateParams
	if err := c.ShouldBindJSON(&contentItemCreateParams); err != nil {
		log.Println("Failed to bind JSON: ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Consumir servicio
	contentItem, err := CreateContentItem(contentItemCreateParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Traducir respuesta
	c.JSON(200, contentItem)
}

func CreateContentItem(contentItemCreateParams domain.ContentItemCreateParams) (contentItem *domain.ContentItem, err error) {

	now := time.Now().UTC()
	contentItem = &domain.ContentItem{
		Category:    contentItemCreateParams.Category,
		Description: contentItemCreateParams.Description,
		Title:       contentItemCreateParams.Title,
		CreatedAt:   &now,
		UpdatedAt:   &now,
	}

	err = InsertContentItem(contentItem)
	if err != nil {
		return nil, fmt.Errorf("error inserting content item: %w", err)
	}

	return contentItem, nil
}

func InsertContentItem(contentItem *domain.ContentItem) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(config.MongoDBURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println("Failed to connect to MongoDB: ", err.Error())
		return fmt.Errorf("error connecting to MongoDB: %w", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
		return fmt.Errorf("error pinging MongoDB: %w", err)
	}

	collection := client.Database(config.MongoDBName).Collection(config.MongoDBCollNameContentItems)

	ctx, cancel = context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	_, err = collection.InsertOne(ctx, contentItem)
	if err != nil {
		log.Println("Failed to insert contentItem to MongoDB: ", err.Error())
		return fmt.Errorf("error inserting contentItem to MongoDB: %w", err)
	}

	return nil
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
