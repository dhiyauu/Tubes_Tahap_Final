package main

import (
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishReportEvent(event TrackingEvent) error {

	body, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = rabbitChannel.Publish(
		"",
		"report_queue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		return err
	}

	log.Println("Berhasil publish ke Report Queue:", string(body))

	return nil
}