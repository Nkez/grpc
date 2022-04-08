package server

import (
	"fmt"
	config "github.com/Nkez/grpc/configs"
	"github.com/Nkez/grpc/internal/service"
	"github.com/Nkez/grpc/pkg/proto"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
)

type ServerGRPC struct {
	server *grpc.Server
}

func NewServerGRPC() *ServerGRPC {
	return &ServerGRPC{grpc.NewServer()}
}

func (s *ServerGRPC) RegisterServices(services *service.Service) {
	proto.RegisterUsersServerServer(s.server, services)
}

func (s *ServerGRPC) Run(cfg *config.Config) error {
	fmt.Println(cfg.GRPC.Port)
	lis, err := net.Listen("tcp", ":"+cfg.GRPC.Port)
	if err != nil {
		log.Fatal().Err(err).Msg("error occurred while running grpc connection")

		return err
	}
	if err := s.server.Serve(lis); err != nil {
		log.Fatal().Err(err).Msg("error occurred while running grpc server")
		return err
	}
	return nil
}

func (s *ServerGRPC) Shutdown() {
	s.server.Stop()
}
