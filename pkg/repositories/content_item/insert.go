package content_item

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jairogloz/go-content-manager/pkg/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func InsertContentItem(contentItem *domain.ContentItem, config domain.EnvVars) error {
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
