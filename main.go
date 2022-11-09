package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"log"
	"test-jenkins/route"
)

func main() {
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

	g := route.Router()

	log.Printf("Server is running at %d port.", viper.GetInt("web.port"))
	log.Fatalf("root err:%s", g.Run(fmt.Sprintf("%s:%d", "0.0.0.0", viper.GetInt("web.port"))))
}
