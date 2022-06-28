package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config") // Register config file name (no extension)
	viper.SetConfigType("json")   // Look for specific type
	viper.ReadInConfig()

	appName := viper.Get("AppName")
	secondKeyName := viper.Get("Keys.key2.Name")
	secondKeyValue := viper.Get("Keys.key2.Value")

	fmt.Printf("The AppName is: %s and the second key has the name: %s and value: %s ", appName, secondKeyName, secondKeyValue)

}
