package main

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"log"
)

func main()  {
	viper.AddConfigPath("config")
	viper.SetConfigName("config-test-jenkins")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("viper ReadInConfig err: %v", err.Error())
	} else {
		all := viper.AllSettings()
		bs, err := yaml.Marshal(all)
		if err != nil {
			log.Fatalf("unable to marshal config to YAML: %v", err.Error())
		}
		log.Println(string(bs))
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("Config file changed: %v", in.Name)
	})
	log.Println("ok")
}
