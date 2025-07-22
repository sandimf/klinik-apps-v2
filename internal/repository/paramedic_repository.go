package repository

import (
	"context"
	"v2/internal/domain"
)

type ParamedicRepository interface {
	Create(ctx context.Context, paramedic *domain.Paramedic) error
}
