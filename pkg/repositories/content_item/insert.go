package content_item

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jairogloz/go-content-manager/pkg/domain"
	"github.com/jairogloz/go-content-manager/pkg/ports"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var _ ports.ContentItemRepository = &Repository{}

type Repository struct {
	config domain.EnvVars
	coll   *mongo.Collection
}

func NewRepository(config domain.EnvVars) (*Repository, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(config.MongoDBURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println("Failed to connect to MongoDB: ", err.Error())
		return nil, fmt.Errorf("error connecting to MongoDB: %w", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
		return nil, fmt.Errorf("error pinging MongoDB: %w", err)
	}

	r := &Repository{
		config: config,
		coll:   client.Database(config.MongoDBName).Collection(config.MongoDBCollNameContentItems),
	}

	return r, nil
}

func (r *Repository) Insert(contentItem *domain.ContentItem) (insertedID string, err error) {

	contentItem.ID = primitive.NewObjectID().Hex()

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	_, err = r.coll.InsertOne(ctx, contentItem)
	if err != nil {
		log.Println("Failed to insert contentItem to MongoDB: ", err.Error())
		return "", fmt.Errorf("error inserting contentItem to MongoDB: %w", err)
	}

	return contentItem.ID, nil
}
