package main

import (
	"fmt"
	"go-template/internal/core/services"
	"go-template/internal/handlers"
	"go-template/internal/repositories"
	"go-template/pkg/config"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func clean(tearDown func()) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		tearDown()
		os.Exit(0)
	}()
}

func main() {
	godotenv.Load(".env")
	// TODO: connect to database please uncommnet below
	// oracledb := datastore.NewOracleDatastore(config.GetOracleConfig())
	// tearDownApp := func() {
	// 	oracledb.CloseConnection()
	// }
	// defer tearDownApp()
	// clean(tearDownApp)

	// perform dependency injection
	messageRepository := repositories.NewMessageRepository(nil) // inject oracledb here
	messageService := services.NewMessageService(messageRepository)
	messageHandler := handlers.NewMessageHandler(messageService)

	// init route with gin
	// can be moved to infrastructure dir to make it cleaner
	router := gin.New()
	router.GET("/message/:id", messageHandler.GetMessage)

	// listen and serve on port that we set in env
	router.Run(fmt.Sprintf(":%s", config.GetAppConfig().Port))
}
