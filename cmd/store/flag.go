package main

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	addrEnv      = "SERVER_ADDRESS"
	addrFlagName = "address"
	addrDefault  = "127.0.0.1:8080"
)

var (
	AddrFlagValue string
)

func FlagsInit() {
	bindEnv()
	bindFlags()
	setFlags()
}

func bindEnv() {
	viper.BindEnv(addrEnv)
}

func bindFlags() {
	pflag.StringP(addrFlagName, "a", addrDefault, "server address")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
}

func setFlags() {
	setAddrFlag()
}

func setAddrFlag() {
	if envValue := viper.GetString(addrEnv); envValue != "" {
		AddrFlagValue = envValue
		return
	}
	AddrFlagValue = viper.GetString(addrFlagName)
}
