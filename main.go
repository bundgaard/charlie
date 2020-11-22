package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

func init() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/charlie")
	viper.AddConfigPath("$HOME/.charlie")
	viper.SetConfigName("charlie")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("readinconfig", err)
	}
}

func main() {
	viper.SetDefault("http.public_folder", "public")
	viper.SetDefault("http.port", "8080")
	viper.SetDefault("http.web_addr", "localhost")

	publicFolder := viper.GetString("http.public_folder")
	webAddr := viper.GetString("http.web_addr")
	port := viper.GetString("http.port")

	root := http.NewServeMux()

	root.Handle("/", http.FileServer(http.Dir(publicFolder)))

	fmt.Printf("Starting server on %s:%s\n", webAddr, port)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", webAddr, port), root); err != nil {
		log.Fatal(err)
	}
}
