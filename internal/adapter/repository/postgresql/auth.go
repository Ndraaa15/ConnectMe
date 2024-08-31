package postgresql

import (
	"gorm.io/gorm"
)

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

type AuthRepository struct {
	db *gorm.DB
}

func (r *AuthRepository) NewAuthRepositoryClient(tx bool) *AuthRepositoryClient {
	if tx {
		return &AuthRepositoryClient{
			q: r.db.Begin(),
		}
	} else {
		return &AuthRepositoryClient{
			q: r.db,
		}
	}
}

type AuthRepositoryClient struct {
	q *gorm.DB
}

func (r *AuthRepositoryClient) Commit() error {
	return r.q.Commit().Error
}

func (r *AuthRepositoryClient) Rollback() error {
	return r.q.Rollback().Error
}
