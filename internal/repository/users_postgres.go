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

func (r *UsersPostgres) GetAllUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User
	query := fmt.Sprintf(`SELECT * from users`)
	err := r.db.Select(&users, query)
	if err != nil {
		log.Error().Err(err).Msg("Err proto restaurant db")
	}
	return users, err
}
