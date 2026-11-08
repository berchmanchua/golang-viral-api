package aws

import (
	"context"
	"fmt"
	"github.com/berchmanchua/golang-viral-api/internal/config"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/aws"
)

func NewDynamoDBClient(appCfg *config.Config) (*dynamodb.Client, error) {
	endpoint := appCfg.DynamoDB.Endpoint

	var opts []func(*awsconfig.LoadOptions) error
	
	opts = append(opts, awsconfig.WithRegion("us-east-1"))
	
	awsCfg, err := awsconfig.LoadDefaultConfig(context.Background(), opts...)
	if err != nil {
			return nil, fmt.Errorf("ddb: load aws config: %w", err)
	}

	ddb := dynamodb.NewFromConfig(awsCfg, func(o *dynamodb.Options) {
			o.BaseEndpoint = aws.String(endpoint)
	})

	return ddb, nil
}
