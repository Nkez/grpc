package service

import (
	"context"
	"github.com/Nkez/grpc/internal/repository"
	"github.com/Nkez/grpc/pkg/proto"
)

type Users interface {
	CreateUser(ctx context.Context, user *proto.User) (*proto.UserId, error)
	GetUserByEmail(ctx context.Context, email *proto.Email) (*proto.User, error)
	GetAllUsers(ctx context.Context, sort *proto.Sort) (*proto.Users, error)
}

type Service struct {
	Users
	proto.UnsafeUsersServerServer
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Users: NewUsersService(repo),
	}
}
