package repository

import (
	"context"
	"github.com/Nkez/grpc/models"
	"github.com/Nkez/grpc/pkg/proto"
	"github.com/jmoiron/sqlx"
)

type Users interface {
	CreateUser(ctx context.Context, user *proto.User) (string, error)
	GetUserByEmail(ctx context.Context, email *proto.Email) (models.User, error)
	SortUsersByStatusAndRegion(ctx context.Context, sort *proto.Sort) ([]models.User, error)
	SortUsersByStatus(ctx context.Context, sort *proto.Sort) ([]models.User, error)
	SortUsersByRegion(ctx context.Context, sort *proto.Sort) ([]models.User, error)
}

type Repository struct {
	Users
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Users: NewUsersPostgres(db),
	}
}
