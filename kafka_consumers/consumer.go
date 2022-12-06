package kafka_consumers

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/wvanbergen/kafka/consumergroup"
	"ic-indexer-worker/app/processor"
	"time"
)

type KafkaConsumerConfig struct {
	Zookeeper     string
	Topic         string
	ConsumerGroup string
}

func InitConsumer(cgroup string, topic string, zookeeperConn string) (*consumergroup.ConsumerGroup, error) {
	// consumer config
	config := consumergroup.NewConfig()
	config.Offsets.Initial = sarama.OffsetOldest
	config.Offsets.ProcessingTimeout = 10 * time.Second

	// join to consumer group
	cg, err := consumergroup.JoinConsumerGroup(cgroup, []string{topic}, []string{zookeeperConn}, config)
	if err != nil {
		return nil, err
	}

	return cg, err
}

func Consume(cg *consumergroup.ConsumerGroup, topic string) {
	for {
		select {
		case msg := <-cg.Messages():
			// messages coming through chanel
			// only take messages from subscribed topic
			if msg.Topic != topic {
				continue
			}


			service := processor.GetIcecreamIndexerWorker()

			err := service.ProcessIcecreamStreamData(msg.Value)

			// commit to zookeeper that message is read
			// this prevent read message multiple times after restart
			err = cg.CommitUpto(msg)
			if err != nil {
				fmt.Println("Error commit zookeeper: ", err.Error())
			}
		}
	}
}
