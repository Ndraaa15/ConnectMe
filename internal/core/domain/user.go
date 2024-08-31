package domain

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID
	FullName string
	Phone    string
	Email    string
	Password string
}
