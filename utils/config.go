package utils

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Configurations exported
type Configurations struct {
	DOL_API_URL_CountryGoods string
	DOL_API_KEY              string
}

// Config provides access to settings
var Config Configurations

func init() {
	loadConfig()
}

/*
loadConfig loads file config details
*/
func loadConfig() {
	// setup
	viper.SetConfigName("config")
	viper.AddConfigPath("./")
	viper.SetConfigType("yml")

	// watch for changes
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})

	// load config
	err := viper.ReadInConfig()
	CheckErr(err)

	err = viper.Unmarshal(&Config)
	CheckErr(err)
}
