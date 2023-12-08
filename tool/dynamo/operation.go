package dynamo

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"reflect"
)

func (c *Client) GetItem(value interface{}) (map[string]*dynamodb.AttributeValue, error) {
	key, err := dynamodbattribute.MarshalMap(value)
	if err != nil {
		return nil, err
	}
	result, err := c.dc.GetItem(&dynamodb.GetItemInput{
		Key:       key,
		TableName: aws.String(c.TableName),
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return result.Item, nil
}

func (c *Client) PutItem(m interface{}) error {
	data, err := dynamodbattribute.MarshalMap(m)
	if err != nil {
		return err
	}

	_, err = c.dc.PutItem(&dynamodb.PutItemInput{
		Item:      data,
		TableName: aws.String(c.TableName),
	})

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteItem(keyName string, keyValue string) error {
	key, err := dynamodbattribute.Marshal(keyValue)
	if err != nil {
		return err
	}

	_, err = c.dc.DeleteItem(&dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			keyName: key,
		},
		TableName: aws.String(c.TableName),
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) ConvertToItem(item map[string]*dynamodb.AttributeValue, v interface{}) error {
	err := dynamodbattribute.UnmarshalMap(item, v)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Scan() ([]map[string]*dynamodb.AttributeValue, error) {
	result, err := c.dc.Scan(&dynamodb.ScanInput{
		TableName: aws.String(c.TableName),
	})
	if err != nil {
		return nil, err
	}
	return result.Items, nil
}

// paginating room, chat
func (c *Client) ScanWithLimit(findKey string, findValue string, limit int32, lastKey interface{}) ([]map[string]*dynamodb.AttributeValue, *string, error) {
	var exclusiveStartKey map[string]*dynamodb.AttributeValue
	if lastKey != nil && !reflect.ValueOf(lastKey).IsNil() {
		key, err := dynamodbattribute.MarshalMap(lastKey)
		if err != nil {
			return nil, nil, err
		}
		exclusiveStartKey = key
	}

	attrValue := dynamodb.AttributeValue{S: aws.String(findValue)}
	input := &dynamodb.QueryInput{
		TableName:              aws.String(c.TableName),
		Limit:                  aws.Int64(int64(limit)),
		ExclusiveStartKey:      exclusiveStartKey,
		KeyConditionExpression: aws.String(fmt.Sprintf("%s = :%s", findKey, findKey)),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			fmt.Sprintf(":%s", findKey): &attrValue,
		},
		ScanIndexForward: aws.Bool(true),
	}

	result, err := c.dc.Query(input)
	if err != nil {
		return nil, nil, err
	}

	var lastEvaluatedKey string
	if result.LastEvaluatedKey != nil {
		lastKeyMap := result.LastEvaluatedKey["id"]
		if lastKeyMap != nil {
			lastEvaluatedKey = *lastKeyMap.S
		}
	}

	return result.Items, &lastEvaluatedKey, nil
}
