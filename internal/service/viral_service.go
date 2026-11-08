package service

import (
	"context"
	"errors"
	"time"

	"github.com/berchmanchua/golang-viral-api/internal/domain"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

const tableName = "ViralItems"

type ViralService interface {
	CreateItem(ctx context.Context, req domain.CreateViralRequest) (*domain.ViralItem, error)
	GetItem(ctx context.Context, id string) (*domain.ViralItem, error)
	UpdateViews(ctx context.Context, id string, newViews int64) error
	DeleteItem(ctx context.Context, id string) error
}

type viralService struct {
	client *dynamodb.Client
}

func NewViralService(client *dynamodb.Client) ViralService {
	return &viralService{
		client: client,
	}
}

func (s *viralService) CreateItem(ctx context.Context, req domain.CreateViralRequest) (*domain.ViralItem, error) {
	if req.Title == "" || req.URL == "" {
		return nil, errors.New("title and url are required")
	}

	item := &domain.ViralItem{
		ID:        uuid.New().String(),
		Title:     req.Title,
		Platform:  req.Platform,
		Views:     0,
		URL:       req.URL,
		CreatedAt: time.Now().UTC(),
	}

	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		return nil, err
	}

	_, err = s.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      av,
	})
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (s *viralService) GetItem(ctx context.Context, id string) (*domain.ViralItem, error) {
	key, err := attributevalue.MarshalMap(map[string]string{"id": id})
	if err != nil {
		return nil, err
	}

	out, err := s.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       key,
	})
	if err != nil {
		return nil, err
	}
	if out.Item == nil {
		return nil, errors.New("item not found")
	}

	var item domain.ViralItem
	err = attributevalue.UnmarshalMap(out.Item, &item)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (s *viralService) UpdateViews(ctx context.Context, id string, newViews int64) error {
	key, err := attributevalue.MarshalMap(map[string]string{"id": id})
	if err != nil {
		return err
	}

	val, err := attributevalue.Marshal(newViews)
	if err != nil {
		return err
	}

	_, err = s.client.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key:       key,
		UpdateExpression: aws.String("SET views = :v"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":v": val,
		},
	})
	
	return err
}

func (s *viralService) DeleteItem(ctx context.Context, id string) error {
	key, err := attributevalue.MarshalMap(map[string]string{"id": id})
	if err != nil {
		return err
	}

	_, err = s.client.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key:       key,
	})
	
	return err
}
