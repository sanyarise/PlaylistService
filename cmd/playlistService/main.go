package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/sanyarise/playlist/config"
	delivery "github.com/sanyarise/playlist/internal/delivery/grpc"
	service "github.com/sanyarise/playlist/internal/delivery/grpc/interface"
	"github.com/sanyarise/playlist/internal/repository"
	"github.com/sanyarise/playlist/internal/usecases"
	"github.com/sanyarise/playlist/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Start load configuration...")
	cfg := config.NewConfig()

	logger := logger.NewLogger(cfg.LogLevel).Logger.Sugar()
	logger.Info("Configuration successfully load")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)

	dns := toDNS(cfg, logger)
	db, err := repository.NewPgxStorage(ctx, dns, logger)
	if err != nil {
		logger.Fatal("can't initialize storage: %v", err)
	}
	store := repository.NewSongRepo(db, logger)

	usecase := usecases.NewSongUsecase(store, logger)
	playlist := usecases.NewPlaylist(logger)

	delivery := delivery.NewDelivery(usecase, playlist, logger)
	server := grpc.NewServer()
	service.RegisterPlaylistServer(server, delivery)

	address := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		logger.Fatal("can't start listener: %v", err)
	}
	go serveGRPC(server, listener)
	logger.Info("Server start on port: %s", cfg.Port)

	<-ctx.Done()
	server.GracefulStop()
	logger.Info("Server stop sucess")

	err = db.ShutDown(cfg.Timeout)
	if err != nil {
		logger.Error(err)
	} else {
		logger.Info("Database connection stopped sucess")
	}
	cancel()
}

func toDNS(cfg *config.Config, l *zap.SugaredLogger) string {
	l.Debug("Enter in main toDNS()")
	dns := fmt.Sprintf("postgres://%s:%s@%s:%s/postgres?sslmode=disable", cfg.PGUser, cfg.PGPass, cfg.PGHost, cfg.PGPort)
	l.Info("dns created success: %s", dns)
	return dns
}

func serveGRPC(server *grpc.Server, listener net.Listener) {
	err := server.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
