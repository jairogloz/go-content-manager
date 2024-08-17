package user

import (
	"context"
	"fmt"
	"time"

	"github.com/jairogloz/go-content-manager/pkg/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (r *Repository) FindByAPIKey(apiKey string) (user *domain.User, error error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	singleResult := r.coll.FindOne(ctx, bson.M{"api_key": apiKey})
	if err := singleResult.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			// No se encontró el documento, pero no hubo error en la consulta
			return nil, nil
		}
		return nil, fmt.Errorf("error finding user: %w", err) // Error finding the document
	}

	if err := singleResult.Decode(&user); err != nil {
		return nil, err // Error decoding the document
	}

	return user, nil
}
