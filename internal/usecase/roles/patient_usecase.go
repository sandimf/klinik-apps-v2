package roles

import (
	"context"
	"v2/internal/domain/roles"
	repo "v2/internal/repository/roles"
)

type PatientUsecase interface {
	CreateOrUpdatePatient(ctx context.Context, patient *roles.Patient) (*roles.Patient, error)
	FindByNIK(ctx context.Context, nik string) (*roles.Patient, error)
	FindAll(ctx context.Context) ([]roles.Patient, error)
	FindAllPaginated(ctx context.Context, page, limit int) ([]roles.Patient, int64, error)
}

type patientUsecase struct {
	repo repo.PatientRepository
}

func NewPatientUsecase(r repo.PatientRepository) PatientUsecase {
	return &patientUsecase{repo: r}
}

func (u *patientUsecase) CreateOrUpdatePatient(ctx context.Context, patient *roles.Patient) (*roles.Patient, error) {
	return u.repo.CreateOrUpdateByNIK(ctx, patient)
}

func (u *patientUsecase) FindByNIK(ctx context.Context, nik string) (*roles.Patient, error) {
	return u.repo.FindByNIK(ctx, nik)
}

func (u *patientUsecase) FindAll(ctx context.Context) ([]roles.Patient, error) {
	return u.repo.FindAll(ctx)
}

func (u *patientUsecase) FindAllPaginated(ctx context.Context, page, limit int) ([]roles.Patient, int64, error) {
	return u.repo.FindAllPaginated(ctx, page, limit)
}
