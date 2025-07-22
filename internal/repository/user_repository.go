package repository

import (
	"context"
	roles "v2/internal/domain/roles"
)

type UserRepository interface {
	Create(ctx context.Context, user *roles.User) (string, error)
	FindByEmail(ctx context.Context, email string) (*roles.User, error)
	FindByID(ctx context.Context, id string) (*roles.User, error)
}
