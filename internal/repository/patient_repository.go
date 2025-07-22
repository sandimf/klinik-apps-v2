package repository

import (
	"context"
	"v2/internal/domain"
)

type PatientRepository interface {
	Create(ctx context.Context, patient *domain.Patient) error
}
