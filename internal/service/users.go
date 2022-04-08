package service

import (
	"context"
	"fmt"
	"github.com/Nkez/grpc/internal/repository"
	"github.com/Nkez/grpc/pkg/proto"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/mail"
)

type UsersService struct {
	repo *repository.Repository
}

func NewUsersService(repo *repository.Repository) *UsersService {
	return &UsersService{repo: repo}
}

func (s *UsersService) CreateUser(ctx context.Context, user *proto.User) (*proto.UserId, error) {
	if _, err := mail.ParseAddress(user.Email); err != nil {
		log.Error().Err(err).Msg("invalid values of fields")
		return &proto.UserId{}, status.Error(400, "invalid values of fields")
	}
	var validate = validator.New()
	if err := validate.Struct(user); err != nil {
		log.Error().Err(err).Msg("invalid values of fields")
		return &proto.UserId{}, status.Error(400, "invalid values of fields")
	}

	id, _ := s.repo.CreateUser(ctx, user)
	fmt.Println(id)
	return &proto.UserId{Id: id}, nil
}

func (s *UsersService) GetUserByEmail(ctx context.Context, email *proto.Email) (*proto.User, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		log.Error().Err(err).Msg(fmt.Sprintf("User with email = %s not found", email.UserEmail))
		return &proto.User{}, status.Error(400, fmt.Sprintf("User with email = %s not found", email.UserEmail))
	}
	return &proto.User{FirstName: user.FirstName, LastName: user.LastName, Email: user.Email, Age: user.Age}, err
}

func (s *UsersService) GetAllUsers(ctx context.Context, empty *emptypb.Empty) (*proto.Users, error) {
	users, err := s.repo.GetAllUsers(ctx)

	protoUsers := make([]*proto.User, 0, len(users))
	for _, val := range users {
		protoUsers = append(protoUsers,
			&proto.User{FirstName: val.FirstName, LastName: val.LastName, Email: val.Email, Age: val.Age})
	}
	return &proto.Users{List: protoUsers}, err
}
