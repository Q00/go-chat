package dynamo

import (
	"fmt"
	"github.com/Q00/go-chat/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var sess *session.Session
var dc *dynamodb.DynamoDB

type Client struct {
	TableName string
	dc        *dynamodb.DynamoDB
}

func GetDynamoSession(cfg *config.AWSInfo) *session.Session {
	if sess == nil {
		accessKey := cfg.AccessKeyID
		secretKey := cfg.SecretKey

		sess = session.Must(session.NewSession(&aws.Config{
			Region:      aws.String(cfg.Region),
			Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
		}))
	}
	return sess
}

func GetDynamoClient(cfg *config.AWSInfo, sess *session.Session, tableName string) *Client {
	if dc == nil {
		dc = dynamodb.New(sess)
		// create table if it doesn't exist
		createTable(cfg, dc)
	}
	return &Client{
		TableName: tableName,
		dc:        dc,
	}
}

// createTable creates a table in DynamoDB first if it doesn't exist
func createTable(cfg *config.AWSInfo, dc *dynamodb.DynamoDB) {
	roomTableInput := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("userId"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("groupId"),
				AttributeType: aws.String("N"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("userId"),
				KeyType:       aws.String("HASH"), // Partition key
			},
			{
				AttributeName: aws.String("groupId"),
				KeyType:       aws.String("RANGE"), // Sort key
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(int64(cfg.Dynamo.Room.ReadCapacityUnits)),
			WriteCapacityUnits: aws.Int64(int64(cfg.Dynamo.Room.WriteCapacityUnits)),
		},
		TableName: aws.String(cfg.Dynamo.Room.TableName),
	}

	_, err := dc.CreateTable(roomTableInput)
	if err != nil {
		fmt.Println(err)
	}

	chatMessageTableInput := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("roomId"),
				AttributeType: aws.String("S"), // Assuming userId is a Number
			},
			{
				AttributeName: aws.String("id"),
				AttributeType: aws.String("S"), // S for String
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("roomId"),
				KeyType:       aws.String("HASH"), // Partition key
			},
			{
				AttributeName: aws.String("id"),
				KeyType:       aws.String("RANGE"), // Sort key
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(int64(cfg.Dynamo.Chat.ReadCapacityUnits)),
			WriteCapacityUnits: aws.Int64(int64(cfg.Dynamo.Chat.WriteCapacityUnits)),
		},
		TableName: aws.String(cfg.Dynamo.Chat.TableName),
	}

	_, err = dc.CreateTable(chatMessageTableInput)
	if err != nil {
		fmt.Println(err)
	}

}
