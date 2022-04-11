package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/Nkez/grpc/internal/repository"
	"github.com/Nkez/grpc/models"
	"github.com/Nkez/grpc/pkg/proto"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/status"
	"net/mail"
	"strings"
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
	if CheckStatus(user.Status) == false {
		return &proto.UserId{}, status.Error(400, "invalid status type")
	}
	if CheckRegion(user.Region) == false {
		return &proto.UserId{}, status.Error(400, "invalid region")
	}
	user.Region = strings.ToUpper(user.Region)
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
	return &proto.User{FirstName: user.FirstName, LastName: user.LastName, Email: user.Email, Age: user.Age, Region: user.Region, Status: user.Status}, err
}

func (s *UsersService) GetAllUsers(ctx context.Context, sort *proto.Sort) (*proto.Users, error) {
	var users []models.User
	err := errors.New(" ")
	if sort.Status != "" && sort.Region != "" {
		users, err = s.repo.SortUsersByStatusAndRegion(ctx, sort)
	}
	if sort.Status != "" && sort.Region == "" {
		users, err = s.repo.SortUsersByStatus(ctx, sort)
	}

	if sort.Status == "" && sort.Region != "" {
		users, err = s.repo.SortUsersByRegion(ctx, sort)
	}

	protoUsers := make([]*proto.User, 0, len(users))
	for _, val := range users {
		protoUsers = append(protoUsers,
			&proto.User{FirstName: val.FirstName, LastName: val.LastName, Email: val.Email, Age: val.Age, Region: val.Region, Status: val.Status})
	}
	return &proto.Users{List: protoUsers}, err
}

func CheckStatus(status string) bool {
	status = strings.ToLower(status)
	switch status {
	case "admin", "customer", "manager":
		return true
	}
	return false
}

func CheckRegion(region string) bool {
	if len(region) != 2 {
		return false
	}
	return true
}
