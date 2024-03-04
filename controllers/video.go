package controllers

import (
	"encoding/json"
	"github/dev-hack95/Textflow/services"
	"github/dev-hack95/Textflow/structs"
	"github/dev-hack95/Textflow/utilities"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func Upload(data structs.VideoPayload) (returnData utilities.ResponseJson) {
	response, err := services.UploadVideo(data)
	if err != nil {
		utilities.ErrorResponse(&returnData, err.Error())
		return
	}

	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "192.168.29.7:9092",
		"acks":              "all",
	})
	if err != nil {
		utilities.ErrorResponse(&returnData, err.Error())
		return
	}

	defer producer.Close()

	topic := "Kafkatopic1"
	deliveryChan := make(chan kafka.Event)

	responseJson, err := json.Marshal(response)
	if err != nil {
		utilities.ErrorResponse(&returnData, err.Error())
		return
	}

	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          responseJson,
	}, deliveryChan)

	if err != nil {
		utilities.ErrorResponse(&returnData, err.Error())
		return
	}

	e := <-deliveryChan
	msg := e.(*kafka.Message)

	if msg.TopicPartition.Error != nil {
		utilities.ErrorResponse(&returnData, msg.TopicPartition.Error.Error())
		return
	}

	utilities.SuccessResponse(&returnData, response)
	return
}
