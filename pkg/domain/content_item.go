package domain

import "time"

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
