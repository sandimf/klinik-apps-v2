package physicalexam

import (
	"context"
	"v2/internal/domain/physicalexam"
)

type PhysicalExaminationRepository interface {
	Create(ctx context.Context, exam *physicalexam.PhysicalExamination) error
	FindByPatientID(ctx context.Context, patientID string) ([]physicalexam.PhysicalExamination, error)
	FindDoctorConsultations(ctx context.Context) ([]physicalexam.PhysicalExamination, error)
	UpdateConsultationStatus(ctx context.Context, id string, status string) error
	Update(ctx context.Context, id string, update map[string]interface{}) error
}
