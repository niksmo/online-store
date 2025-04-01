package main

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	kafkaServersEnv      = "KAFKA_SERVERS"
	kafkaServersFlagName = "kafka-servers"
	kafkaServersDefault  = "kafka-1:9094,kafka-2:9094"

	kafkaTopicEnv      = "KAFKA_TOPIC"
	kafkaTopicFlagName = "kafka-topic"
	kafkaTopicDefault  = "orders_create"

	kafkaConsumerGroupEnv      = "KAFKA_CONSUMER_GROUP"
	kafkaConsumerGroupFlagName = "kafka-consumer-group"
	kafkaConsumerGroupDefault  = "1"
)

var (
	KafkaServersFlagValue       string
	KafkaTopicFlagValue         string
	KafkaConsumerGroupFlagValue string
)

func FlagsInit() {
	bindEnv()
	bindFlags()
	setFlags()
}

func bindEnv() {
	viper.BindEnv(kafkaServersEnv)
	viper.BindEnv(kafkaTopicEnv)
	viper.BindEnv(kafkaConsumerGroupEnv)
}

func bindFlags() {
	pflag.StringP(
		kafkaServersFlagName,
		"b",
		kafkaServersDefault,
		"kafka bootstrap servers",
	)

	pflag.StringP(
		kafkaTopicFlagName,
		"t",
		kafkaTopicDefault,
		"kafka topic",
	)

	pflag.StringP(
		kafkaConsumerGroupFlagName,
		"g",
		kafkaConsumerGroupDefault,
		"kafka consumer group",
	)

	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}

func setFlags() {
	setKafkaServersFlag()
	setKafkaTopicFlag()
	setKafkaConsumerGroupFlag()
}

func setKafkaServersFlag() {
	if envValue := viper.GetString(kafkaServersEnv); envValue != "" {
		KafkaServersFlagValue = envValue
		return
	}
	KafkaServersFlagValue = viper.GetString(kafkaServersFlagName)
}

func setKafkaTopicFlag() {
	if envValue := viper.GetString(kafkaTopicEnv); envValue != "" {
		KafkaTopicFlagValue = envValue
		return
	}
	KafkaTopicFlagValue = viper.GetString(kafkaTopicFlagName)
}

func setKafkaConsumerGroupFlag() {
	if envValue := viper.GetString(kafkaConsumerGroupEnv); envValue != "" {
		KafkaConsumerGroupFlagValue = envValue
		return
	}
	KafkaConsumerGroupFlagValue = viper.GetString(kafkaConsumerGroupFlagName)
}
