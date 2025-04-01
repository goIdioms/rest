package main

import (
	"log"

	"github.com/goIdioms/rest"
	"github.com/goIdioms/rest/pkg/handlers"
	"github.com/goIdioms/rest/pkg/repository"
	"github.com/goIdioms/rest/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	initConfig()
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handlers.NewHandler(services)

	server := new(rest.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error running http server %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}
