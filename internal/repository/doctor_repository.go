package repository

import (
	"context"
	"v2/internal/domain"
)

type DoctorRepository interface {
	Create(ctx context.Context, doctor *domain.Doctor) error
}
