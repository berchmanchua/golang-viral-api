package domain

import "time"

type ViralItem struct {
	ID          string    `json:"id" dynamodbav:"id"`
	Title       string    `json:"title" dynamodbav:"title"`
	Description string    `json:"description" dynamodbav:"description"`
	Platform    string    `json:"platform" dynamodbav:"platform"`
	Views       int64     `json:"views" dynamodbav:"views"`
	URL         string    `json:"url" dynamodbav:"url"`
	CreatedAt   time.Time `json:"created_at" dynamodbav:"created_at"`
}

type CreateViralRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Platform    string `json:"platform"`
	URL         string `json:"url"`
}
