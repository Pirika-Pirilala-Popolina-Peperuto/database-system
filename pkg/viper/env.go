package viper

import (
	"strings"

	"github.com/spf13/viper"
)

func AutoEnv() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}
