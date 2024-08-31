package port

import "github.com/Ndraaa15/ConnectMe/internal/adapter/repository/postgresql"

type AuthRepositoryItf interface {
	NewAuthRepositoryClient(tx bool) *postgresql.AuthRepositoryClient
}

type AuthRepositoryClientItf interface {
	Commit() error
	Rollback() error
}

type AuthServiceItf interface{}
