package main

import (
	"ecommerceuser/middleware"
	grpc_api "ecommerceuser/proto"
	grpcapis "ecommerceuser/user/gRPCAPIs"
	"log"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.SetTrustedProxies([]string{"localhost"})
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Hello World"})
	})
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	apiGroup := r.Group("/api/v1")
	apiGroup.Use(middleware.Logger())
	{
		addRoutes(apiGroup)
	}

	return r
}

func startGRPCServer(lis net.Listener) {
	s := grpc.NewServer()
	api := grpcapis.NewGrpcServer()
	grpc_api.RegisterUserMessengerServer(s, api)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	lis := setUpAndRunGRPCServer()
	r := setupRouter()
	go startGRPCServer(lis)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}

func setUpAndRunGRPCServer() net.Listener {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	return lis
}
