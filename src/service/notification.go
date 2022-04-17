package service

import (
	"context"
	"customer-api/src/contract"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type notificationService struct {
	config   aws.Config
	topicARN string
}

func NewNotificationService(config aws.Config, topicARN string) contract.NotificationService {
	return notificationService{
		config:   config,
		topicARN: topicARN,
	}
}

func (s notificationService) Publish(ctx context.Context, message string) error {
	snsClient := sns.NewFromConfig(s.config)

	publishInput := &sns.PublishInput{
		TopicArn: aws.String(s.topicARN),
		Message:  aws.String(message),
	}

	_, err := snsClient.Publish(ctx, publishInput)
	return err
}
