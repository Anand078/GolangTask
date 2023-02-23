package config

import (
	"MS1/constants"
	"MS1/model"
	"context"
	"encoding/json"
	"strconv"

	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

type Publisher struct {
	Writer      *kafka.Writer
	MessagePool chan *model.Product
}

func InitKafka() (publisher *Publisher) {
	log.Infoln("InitKafka() function started.....")
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{constants.KafkaBrokerAddress},
		Topic:   constants.KafkaTopic,
	})

	workerPool := make(chan *model.Product)
	pb := &Publisher{Writer: w, MessagePool: workerPool}
	go pb.pullMessageAndPublish()
	log.Infoln("InitKafka() function ended.....")
	return pb
}

func (p *Publisher) PublishKafka(product *model.Product) { //pass here model
	p.MessagePool <- product
}

func (p *Publisher) pullMessageAndPublish() {
	for elem := range p.MessagePool {
		productByte, err := json.Marshal(elem)
		if err != nil {
			log.Infoln(err)
			continue
		}
		p.Writer.WriteMessages(context.Background(), kafka.Message{Key: []byte(strconv.Itoa(int(elem.ID))), Value: productByte})
	}
}
