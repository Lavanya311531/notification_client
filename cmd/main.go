package main

import (
	"context"

	"notification_client/internal/config"
	"notification_client/internal/models"
	"notification_client/internal/producer"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

func main() {

	config.LoadEnv()

	client, err := azservicebus.NewClientFromConnectionString(
		config.GetServiceBusConnectionString(),
		nil,
	)

	if err != nil {
		panic(err)
	}

	sender, err := client.NewSender(
		config.GetQueueName(),
		nil,
	)

	if err != nil {
		panic(err)
	}

	defer sender.Close(context.Background())

	prod := producer.NewProducer(sender)

	payload := models.Notification{
		NotificationID: "NOTIF-005",
		Recipient:      "lavu8599@gmail.com",
		Channel:        "EMAIL",
		Message:        "Lavanya R",
	}

	err = prod.Send(payload)

	if err != nil {
		panic(err)
	}

	println("Notification Published Successfully")
	println("Notification ID:", payload.NotificationID)
	println("Message:", payload.Message)
}
