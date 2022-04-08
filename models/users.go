package models

type User struct {
	ID        string `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"first_name" validate:"required"`
	LastName  string `json:"last_name" db:"last_name" validate:"required"`
	Email     string `json:"email" db:"email"  validate:"email" `
	Age       int64  `json:"age" db:"age" validate:"required,alpha"`
}
