package domain

import "time"

// ContentItemCreateParams is the struct that represents the request body for creating a content item
type ContentItemCreateParams struct {
	Category    string `json:"category"`
	Description string `json:"description"`
	Title       string `json:"title"`
}

// ContentItemUpdateParams is the struct that represents the request body for updating a content item
type ContentItemUpdateParams struct {
	Category              *string    `json:"category" bson:"category,omitempty"`
	Description           *string    `json:"description" bson:"description,omitempty"`
	PublishedAtBlog       *time.Time `json:"published_at_blog" bson:"published_at_blog,omitempty"`
	PublishedAtLinkedIn   *time.Time `json:"published_at_linkedin" bson:"published_at_linkedin,omitempty"`
	PublishedAtNewsletter *time.Time `json:"published_at_newsletter" bson:"published_at_newsletter,omitempty"`
	PublishedAtYoutube    *time.Time `json:"published_at_youtube" bson:"published_at_youtube,omitempty"`
	Title                 *string    `json:"title" bson:"title,omitempty"`
	UpdatedAt             *time.Time `json:"-" bson:"updated_at,omitempty"`
}

// ContentItem reflects a content item in the system.
type ContentItem struct {
	ID                    string     `json:"id" bson:"_id"`
	Category              string     `json:"category" bson:"category"`
	CreatedAt             *time.Time `json:"created_at" bson:"created_at"`
	Description           string     `json:"description" bson:"description"`
	PublishedAtBlog       *time.Time `json:"published_at_blog" bson:"published_at_blog"`
	PublishedAtLinkedIn   *time.Time `json:"published_at_linkedin" bson:"published_at_linkedin"`
	PublishedAtNewsletter *time.Time `json:"published_at_newsletter" bson:"published_at_newsletter"`
	PublishedAtYoutube    *time.Time `json:"published_at_youtube" bson:"published_at_youtube"`
	Title                 string     `json:"title" bson:"title"`
	UpdatedAt             *time.Time `json:"updated_at" bson:"updated_at"`
	UserID                string     `json:"-" bson:"user_id"`
}

// ContentItemListResponse is the struct that represents the response body for listing content items
type ContentItemListResponse struct {
	ContentItems []*ContentItem `json:"content_items"`
	TotalCount   int            `json:"total_count"`
	Count        int            `json:"count"`
	Page         int            `json:"page"`
	TotalPages   int            `json:"total_pages"`
}
