package producer

import (
	"context"
	"encoding/json"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

type Producer struct {
	Sender *azservicebus.Sender
}

func NewProducer(sender *azservicebus.Sender) *Producer {
	return &Producer{
		Sender: sender,
	}
}

func (p *Producer) Send(data interface{}) error {

	body, err := json.Marshal(data)

	if err != nil {
		return err
	}

	msg := &azservicebus.Message{
		Body: body,
	}

	return p.Sender.SendMessage(
		context.Background(),
		msg,
		nil,
	)
}
