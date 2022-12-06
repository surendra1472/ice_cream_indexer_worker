package config

import (
	"fmt"
	"github.com/spf13/viper"
	"ic-indexer-worker/kafka_consumers"
	"os"
	"strings"
)

var config *Config


type Config struct {
	Server  ServerConfig
	Kafka         KafkaConfiguration
}

type ServerConfig struct {
	Port string
}


type KafkaConfiguration struct {
	Brokers             []string
	IceCreamCreateTopic string
	Zookeeper           string
	ConsumerGroup       string
}


func GetConfig() *Config {
	return config
}


func InitializeKafkaConsumer() {
	consumerConfig := kafka_consumers.KafkaConsumerConfig{}
	consumerConfig.Zookeeper = GetConfig().Kafka.Zookeeper
	consumerConfig.Topic = GetConfig().Kafka.IceCreamCreateTopic
	consumerConfig.ConsumerGroup = GetConfig().Kafka.ConsumerGroup
	cg, err := kafka_consumers.InitConsumer(consumerConfig.ConsumerGroup, consumerConfig.Topic, consumerConfig.Zookeeper)
	if err != nil {
		fmt.Println("Error consumer goup: ", err.Error())
		os.Exit(1)
	}
	defer cg.Close()

	// run consumer
	kafka_consumers.Consume(cg, consumerConfig.Topic)
}

func Initialize() error {
	viper.SetConfigFile(getConfigFile())
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	for _, k := range viper.AllKeys() {
		value := viper.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			viper.Set(k, os.Getenv((strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}"))))
		}
	}
	return viper.Unmarshal(&config)
}


func getConfigFile() string {
	return "app/config/" + os.Getenv("ENV") + ".yml"
}
