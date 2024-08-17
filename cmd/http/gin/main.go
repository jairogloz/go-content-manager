package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	content_item_hdlr "github.com/jairogloz/go-content-manager/cmd/http/gin/handlers/content_item"
	"github.com/jairogloz/go-content-manager/pkg/domain"
	contentItemRepo "github.com/jairogloz/go-content-manager/pkg/repositories/content_item"
	userRepo "github.com/jairogloz/go-content-manager/pkg/repositories/user"
	"github.com/jairogloz/go-content-manager/pkg/services/content_item"
	userService "github.com/jairogloz/go-content-manager/pkg/services/user"
	"github.com/spf13/viper"
)

var config domain.EnvVars

func main() {
	var err error
	config, err = LoadConfig()
	if err != nil {
		log.Fatalf("error loading config: %s", err.Error())
	}

	r := gin.Default()

	userRepo, err := userRepo.NewRepository(config)
	if err != nil {
		log.Fatalf("error creating user repository: %s", err.Error())
	}

	userService, err := userService.NewService(userRepo)
	if err != nil {
		log.Fatalf("error creating user service: %s", err.Error())
	}

	r.Use(content_item_hdlr.AuthMiddleware(userService))

	repo, err := contentItemRepo.NewRepository(config)
	if err != nil {
		log.Fatalf("error creating content item repository: %s", err.Error())
	}

	service, err := content_item.NewService(repo)
	if err != nil {
		log.Fatalf("error creating content item service: %s", err.Error())
	}

	contentItemHdlr, err := content_item_hdlr.NewHttpHandler(service, config)
	if err != nil {
		log.Fatalf("error creating content item handler: %s", err.Error())
	}

	r.POST("/content", contentItemHdlr.Create)
	r.GET("/content/:id", contentItemHdlr.Get)
	r.GET("/content/", contentItemHdlr.List)
	r.DELETE("/content/:id", contentItemHdlr.Delete)
	r.PUT("/content/:id", contentItemHdlr.Update)

	log.Fatalln(r.Run(fmt.Sprintf(":%s", config.ServerPort))) // listen and serve on 0.0.0.0:8080 by default
}

func LoadConfig() (config domain.EnvVars, err error) {

	viper.AddConfigPath("../../../.")
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	err = viper.ReadInConfig()
	if err != nil {
		return config, fmt.Errorf("error reading config file: %w", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, fmt.Errorf("error unmarshalling config: %w", err)
	}

	return config, nil

}
