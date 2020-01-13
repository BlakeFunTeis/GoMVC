package main

import (
	"GoMVC/core"
	"GoMVC/routes"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)


func main() {
	ginInstance := gin.New()
	ginInstance.Use(gin.Logger())
	ginInstance.Use(gin.Recovery())
	gin.SetMode(getMode())
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()
	defer core.CloseDatabase()
	defer core.CloseElasticConnect()
	defer core.RedisClose()
	routesInstance := routes.RouterInstance(ginInstance)
	ReadTimeOut, _ := time.ParseDuration(os.Getenv("read_timeout"))
	WriteTimeOut, _ := time.ParseDuration(os.Getenv("write_timeout"))
	server := &http.Server{
		Addr:           ":" + os.Getenv("http_port"),
		Handler:        routesInstance,
		ReadTimeout:    ReadTimeOut,
		WriteTimeout:   WriteTimeOut,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}

func getMode() string {
	var mode = os.Getenv("app_env")
	switch strings.ToLower(mode) {
	case "debug":
		return gin.DebugMode
	case "local":
		return gin.DebugMode
	case "release":
		return gin.ReleaseMode
	case "test":
		return gin.TestMode
	default:
		return gin.DebugMode
	}
}
