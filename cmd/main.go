package main

import (
	config "github.com/Nkez/grpc/configs"
	"github.com/Nkez/grpc/internal/repository"
	"github.com/Nkez/grpc/internal/service"
	"github.com/Nkez/grpc/server"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.Init("./configs")
	if err != nil {
		log.Fatal().Err(err).Msg("wrong config variables")
	}

	db, err := newPostgresDB(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("err initializing DB")
	}

	repo := repository.NewRepository(db)
	servicesGRPC := service.NewService(repo)
	srvGRPC := server.NewServerGRPC()
	srvGRPC.RegisterServices(servicesGRPC)

	go func() {
		if err := srvGRPC.Run(cfg); err != nil {
			log.Error().Err(err).Msg("error occurred while running gRPC server")
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	srvGRPC.Shutdown()

	if err := db.Close(); err != nil {
		log.Fatal().Err(err).Msg("failed to stop connection db")
	}
}

func newPostgresDB(cfg *config.Config) (*sqlx.DB, error) {
	return repository.NewPostgresDB(repository.Config{
		Host:     cfg.Postgres.Host,
		Port:     cfg.Postgres.Port,
		Username: cfg.Postgres.Username,
		Password: cfg.Postgres.Password,
		DBName:   cfg.Postgres.Dbname,
		SSLMode:  cfg.Postgres.Sslmode,
	})
}
