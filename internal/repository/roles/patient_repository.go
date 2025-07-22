package roles

import (
	"context"
	"v2/internal/domain/roles"
)

type PatientRepository interface {
	Create(ctx context.Context, patient *roles.Patient) error
	CreateOrUpdateByNIK(ctx context.Context, patient *roles.Patient) (*roles.Patient, error)
	FindByNIK(ctx context.Context, nik string) (*roles.Patient, error)
	FindAll(ctx context.Context) ([]roles.Patient, error)
	FindAllPaginated(ctx context.Context, page, limit int) ([]roles.Patient, int64, error)
}
