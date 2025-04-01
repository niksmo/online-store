package main

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	addrEnv      = "SERVER_ADDRESS"
	addrFlagName = "address"
	addrDefault  = ":8000"

	kafkaServersEnv      = "KAFKA_SERVERS"
	kafkaServersFlagName = "kafka-servers"
	kafkaServersDefault  = "kafka-1:9094,kafka-2:9094"

	kafkaTopicEnv      = "KAFKA_TOPIC"
	kafkaTopicFlagName = "kafka-topic"
	kafkaTopicDefault  = "orders_create"
)

var (
	AddrFlagValue         string
	KafkaServersFlagValue string
	KafkaTopicFlagValue   string
)

func FlagsInit() {
	bindEnv()
	bindFlags()
	setFlags()
}

func bindEnv() {
	viper.BindEnv(addrEnv)
	viper.BindEnv(kafkaServersEnv)
	viper.BindEnv(kafkaTopicEnv)
}

func bindFlags() {
	pflag.StringP(addrFlagName, "a", addrDefault, "server address")

	pflag.StringP(
		kafkaServersFlagName, "b", kafkaServersDefault, "kafka bootstrap servers",
	)

	pflag.StringP(kafkaTopicFlagName, "t", kafkaTopicDefault, "kafka topic")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}

func setFlags() {
	setAddrFlag()
	setKafkaServersFlag()
	setKafkaTopicFlag()
}

func setAddrFlag() {
	if envValue := viper.GetString(addrEnv); envValue != "" {
		AddrFlagValue = envValue
		return
	}
	AddrFlagValue = viper.GetString(addrFlagName)
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
