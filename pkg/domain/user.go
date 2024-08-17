package domain

// User represents a user in the system.
type User struct {
	APIKey string `bson:"api_key"`
	Email  string `bson:"email"`
	Name   string `bson:"name"`
	UserID string `bson:"user_id"`
}
