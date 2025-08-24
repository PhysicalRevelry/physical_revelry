package services

import (
    "context"
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DynamoDBService struct {
    client *dynamodb.Client
    tableName string
}

func NewDynamoDBService(tableName string) (*DynamoDBService, error) {
    cfg, err := config.LoadDefaultConfig(context.TODO(),
        config.WithRegion("us-east-1"),
        config.WithEndpointResolver(aws.EndpointResolverFunc(
            func(service, region string) (aws.Endpoint, error) {
                return aws.Endpoint{URL: "http://localhost:8000"}, nil
            })),
    )
    if err != nil {
        return nil, err
    }

    client := dynamodb.NewFromConfig(cfg)

    return &DynamoDBService{
        client: client,
        tableName: tableName,
    }, nil
}

func (d *DynamoDBService) CreateTable(ctx context.Context) error {
    _, err := d.client.CreateTable(ctx, &dynamodb.CreateTableInput{
        TableName: aws.String(d.tableName),
        KeySchema: []types.KeySchemaElement{
            {
                AttributeName: aws.String("PK"),
                KeyType:       types.KeyTypeHash,
            },
            {
                AttributeName: aws.String("SK"),
                KeyType:       types.KeyTypeRange,
            },
        },
        AttributeDefinitions: []types.AttributeDefinition{
            {
                AttributeName: aws.String("PK"),
                AttributeType: types.ScalarAttributeTypeS,
            },
            {
                AttributeName: aws.String("SK"),
                AttributeType: types.ScalarAttributeTypeS,
            },
        },
        BillingMode: types.BillingModePayPerRequest,
    })

    return err
}