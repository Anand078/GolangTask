package config

import (
	"MS1/constants"
	"MS1/model"
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/segmentio/kafka-go"
)

type Publisher struct {
	Writer      *kafka.Writer
	MessagePool chan *model.Product
}

func InitKafka() (publisher *Publisher) {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{constants.KafkaBrokerAddress},
		Topic:   constants.KafkaTopic,
	})

	workerPool := make(chan *model.Product)
	pb := &Publisher{Writer: w, MessagePool: workerPool}
	go pb.pullMessageAndPublish()
	return pb
}

func (p *Publisher) PublishKafka(product *model.Product) { //pass here model
	p.MessagePool <- product
}

func (p *Publisher) pullMessageAndPublish() {
	for elem := range p.MessagePool {
		productByte, err := json.Marshal(elem)
		if err != nil {
			fmt.Println(err)
			continue
		}
		p.Writer.WriteMessages(context.Background(), kafka.Message{Key: []byte(strconv.Itoa(int(elem.ID))), Value: productByte})
	}
}
