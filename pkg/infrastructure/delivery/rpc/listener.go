package rpc

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
)

type MessageListener struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      amqp.Queue
	handler    *Handler
}

func NewMessageListener(handler *Handler) MessageListener {
	url := "amqp://guest:guest@rabbit:5672"

	connection := connect(url)
	channel := openChannel(connection)

	return MessageListener{
		connection: connection,
		channel:    channel,
		handler:    handler,
	}
}

func connect(url string) *amqp.Connection {
	connection, err := amqp.Dial(url)

	if err != nil {
		panic("could not establish connection with RabbitMQ:" + err.Error())
	}

	return connection
}

func openChannel(connection *amqp.Connection) *amqp.Channel {
	channel, err := connection.Channel()

	if err != nil {
		panic("could not open RabbitMQ channel:" + err.Error())
	}

	err = channel.ExchangeDeclare("fmf-menu", "fanout", true, false, false, false, nil)

	if err != nil {
		panic(err)
	}

	_, err = channel.QueueDeclare("fmf-menu", true, false, false, false, nil)

	if err != nil {
		panic("error declaring the queue: " + err.Error())
	}

	err = channel.QueueBind("fmf-menu", "#", "fmf-menu", false, nil)

	if err != nil {
		panic("error binding to the queue: " + err.Error())
	}

	return channel
}

func (ml MessageListener) Listen() {
	msgs, err := ml.channel.Consume("fmf-menu", "", false, false, false, false, nil)

	if err != nil {
		panic("error consuming the queue: " + err.Error())
	}

	for msg := range msgs {
		go ml.processMessage(msg)
		msg.Ack(false)
	}
}

func (ml MessageListener) processMessage(message amqp.Delivery) {
	response, err := ml.handler.Process(string(message.Body))

	if err != nil {
		dat, _ := json.Marshal(map[string]string{ "error": err.Error()})
		ml.respond(message, string(dat))
	}

	if response != "" {
		ml.respond(message, response)
	}
}

func  (ml MessageListener) respond(message amqp.Delivery, response string) {
	err := ml.channel.Publish(
		"",
		message.ReplyTo,
		false,
		false,
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: message.CorrelationId,
			Body:          []byte(response),
		},
	)

	if err != nil {
		fmt.Println(err.Error())
	}
}
