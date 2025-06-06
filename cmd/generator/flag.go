package main

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	workersEnv      = "N_WORKERS"
	workersFlagName = "workers"
	workersDefault  = 8

	addrEnv      = "URL_ADDRESS"
	addrFlagName = "url"
	addrDefault  = "http://store:8000/"
)

var (
	WorkersFlagValue int
	AddrFlagValue    string
)

func FlagsInit() {
	bindEnv()
	bindFlags()
	setFlags()
}

func bindEnv() {
	viper.BindEnv(workersEnv)
	viper.BindEnv(addrEnv)
}

func bindFlags() {
	pflag.IntP(workersFlagName, "w", workersDefault, "number of order senders")
	pflag.StringP(addrFlagName, "a", addrDefault, "server URL address")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}

func setFlags() {
	setWorkersFlag()
	setAddrFlag()
}

func setWorkersFlag() {
	if envValue := viper.GetInt(workersEnv); envValue != 0 {
		WorkersFlagValue = envValue
		return
	}
	WorkersFlagValue = viper.GetInt(workersFlagName)
}

func setAddrFlag() {
	if envValue := viper.GetString(addrEnv); envValue != "" {
		AddrFlagValue = envValue
		return
	}
	AddrFlagValue = viper.GetString(addrFlagName)
}
