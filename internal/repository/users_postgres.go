package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Nkez/grpc/models"
	"github.com/Nkez/grpc/pkg/proto"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type UsersPostgres struct {
	db *sqlx.DB
}

func NewUsersPostgres(db *sqlx.DB) *UsersPostgres {
	return &UsersPostgres{db: db}
}

func (r *UsersPostgres) CreateUser(ctx context.Context, user *proto.User) (string, error) {
	id := ""
	query := fmt.Sprintf("INSERT INTO users (first_name, last_name, email, age, region, status) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id")
	if err := r.db.QueryRow(query, user.FirstName, user.LastName, user.Email, user.Age, user.Region, user.Status).Scan(&id); err != nil {
		log.Error().Err(err).Msg("")
	}
	return id, nil
}

func (r *UsersPostgres) GetUserByEmail(ctx context.Context, email *proto.Email) (models.User, error) {

	var user models.User
	query := fmt.Sprintf("select first_name,last_name , email , age  from users where email = $1")
	if err := r.db.QueryRow(query,
		email.UserEmail,
	).Scan(
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Age,
	); err != nil {
		if err == sql.ErrNoRows {
			return user, err
		}
		return user, err
	}
	return user, nil
}

func (r *UsersPostgres) SortUsersByStatusAndRegion(ctx context.Context, sort *proto.Sort) ([]models.User, error) {
	var users []models.User
	query := fmt.Sprintf(`select u.first_name, u.last_name, u.email, u.age, u.region , status  from  users u 
									where u.status = $1 and u.region = $2`)
	err := r.db.Select(&users, query, sort.Status, sort.Region)
	if err != nil {
		log.Error().Err(err).Msg("Err user sort by status and region db")
	}
	return users, err
}

func (r *UsersPostgres) SortUsersByStatus(ctx context.Context, sort *proto.Sort) ([]models.User, error) {
	var users []models.User
	query := fmt.Sprintf(`select u.first_name, u.last_name, u.email, u.age, u.region , status  from  users u 
									where u.status = $1`)
	err := r.db.Select(&users, query, sort.Status)
	if err != nil {
		log.Error().Err(err).Msg("Err user sort by status db")
	}
	return users, err
}

func (r *UsersPostgres) SortUsersByRegion(ctx context.Context, sort *proto.Sort) ([]models.User, error) {
	var users []models.User
	query := fmt.Sprintf(`select u.first_name, u.last_name, u.email, u.age, u.region , status  from  users u 
									where u.region = $1`)
	err := r.db.Select(&users, query, sort.Region)
	if err != nil {
		log.Error().Err(err).Msg("Err user sort by status db")
	}
	return users, err
}
