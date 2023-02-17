package main

import (
	"context"
	"fmt"
	"net"

	"gitlab.udevs.io/template/template_notification_service/grpc/client"
	"gitlab.udevs.io/template/template_notification_service/storage/postgres"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gitlab.udevs.io/template/template_notification_service/config"
	"gitlab.udevs.io/template/template_notification_service/grpc"
	"gitlab.udevs.io/template/template_notification_service/pkg/logger"
)

func main() {
	// Load from .env file to Operating System
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("No .env file found")
	}

	cfg := config.Load()

	log := logger.New(cfg.LogLevel, cfg.ServiceName)
	defer logger.Cleanup(log)

	log.Info("main: pgxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase),
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pgStore, err := postgres.NewPostgres(ctx, cfg, log)
	if err != nil {
		log.Fatal("postgres.NewPostgres", logger.Error(err))
	}
	defer pgStore.CloseDB()

	scvs, err := client.New(ctx, cfg)
	if err != nil {
		log.Fatal("client.NewGrpcClients: %v", logger.Error(err))
	}

	// @TODO:: move websocket part to cmd/main
	grpcServer := grpc.SetUpServer(cfg, log, pgStore, scvs)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s", cfg.RPCPort))
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	log.Info("main: server running", logger.String("port", cfg.RPCPort))
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	log.Info("http server running on port 8080")
}
