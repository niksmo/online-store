package main

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	kafkaServersEnv      = "KAFKA_SERVERS"
	kafkaServersFlagName = "kafka-servers"
	kafkaServersDefault  = "127.0.0.1:19094,127.0.0.1:29094"

	kafkaTopicEnv      = "KAFKA_TOPIC"
	kafkaTopicFlagName = "kafka-topic"
	kafkaTopicDefault  = "orders_create"
)

var (
	KafkaServersFlagValue string
	KafkaTopicFlagValue   string
)

func FlagsInit() {
	bindEnv()
	bindFlags()
	setFlags()
}

func bindEnv() {
	viper.BindEnv(kafkaServersEnv)
	viper.BindEnv(kafkaTopicEnv)
}

func bindFlags() {
	pflag.StringP(kafkaServersFlagName, "b", kafkaServersDefault, "kafka bootstrap servers")
	pflag.StringP(kafkaTopicFlagName, "t", kafkaTopicDefault, "kafka topic")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}

func setFlags() {
	setKafkaServersFlag()
	setKafkaTopicFlag()
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
