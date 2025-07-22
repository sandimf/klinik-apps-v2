package repository

import (
	"context"
	"v2/internal/domain"
)

type AdminRepository interface {
	Create(ctx context.Context, admin *domain.Admin) error
}
