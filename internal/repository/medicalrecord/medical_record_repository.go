package medicalrecord

import (
	"context"
	"v2/internal/domain/medicalrecord"

	"github.com/google/uuid"
)

type MedicalRecordRepository interface {
	Create(ctx context.Context, mr *medicalrecord.MedicalRecord) error
	FindByPatientID(ctx context.Context, patientID uuid.UUID) (*medicalrecord.MedicalRecord, error)
	FindByMRNumber(ctx context.Context, mrNumber string) (*medicalrecord.MedicalRecord, error)
}
