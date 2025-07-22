package physicalexam

import (
	"context"
	"v2/internal/domain/physicalexam"
	repo "v2/internal/repository/physicalexam"
)

type PhysicalExaminationUsecase interface {
	Create(ctx context.Context, exam *physicalexam.PhysicalExamination) error
	FindByPatientID(ctx context.Context, patientID string) ([]physicalexam.PhysicalExamination, error)
	FindDoctorConsultations(ctx context.Context) ([]physicalexam.PhysicalExamination, error)
	UpdateConsultationStatus(ctx context.Context, id string, status string) error
	Update(ctx context.Context, id string, update map[string]interface{}) error
}

type physicalExaminationUsecase struct {
	repo repo.PhysicalExaminationRepository
}

func NewPhysicalExaminationUsecase(r repo.PhysicalExaminationRepository) PhysicalExaminationUsecase {
	return &physicalExaminationUsecase{repo: r}
}

func (u *physicalExaminationUsecase) Create(ctx context.Context, exam *physicalexam.PhysicalExamination) error {
	return u.repo.Create(ctx, exam)
}

func (u *physicalExaminationUsecase) FindByPatientID(ctx context.Context, patientID string) ([]physicalexam.PhysicalExamination, error) {
	return u.repo.FindByPatientID(ctx, patientID)
}

func (u *physicalExaminationUsecase) FindDoctorConsultations(ctx context.Context) ([]physicalexam.PhysicalExamination, error) {
	return u.repo.FindDoctorConsultations(ctx)
}

func (u *physicalExaminationUsecase) UpdateConsultationStatus(ctx context.Context, id string, status string) error {
	return u.repo.UpdateConsultationStatus(ctx, id, status)
}

func (u *physicalExaminationUsecase) Update(ctx context.Context, id string, update map[string]interface{}) error {
	return u.repo.Update(ctx, id, update)
}
