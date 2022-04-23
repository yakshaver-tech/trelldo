package main

import (
	"github.com/spf13/viper"
	"github.com/yakshaver-tech/trelldo/cmd"
)

func main() {
	cmd.Execute()
	viper.WriteConfig()
}
