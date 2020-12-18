package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"sl_go_notify/sqs/slack"
)

func Handler(ctx context.Context, sqsEvent events.SQSEvent) error {

	channelName := "#kobayashi-dev"

	for _, message := range sqsEvent.Records {
		fmt.Printf("message %s ,source %s = %s \n", message.MessageId, message.EventSource, message.Body)

		err := slack.Notify(message.Body, channelName)
		if err != nil {
			panic(err)
		}
	}

	return nil
}

func main() {
	lambda.Start(Handler)
}
