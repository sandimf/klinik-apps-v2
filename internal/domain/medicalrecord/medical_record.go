package medicalrecord

import (
	"time"

	"github.com/google/uuid"
)

type MedicalRecord struct {
	ID        uuid.UUID `json:"id"`
	PatientID uuid.UUID `json:"patient_id"`
	MRNumber  string    `json:"mr_number"`
	CreatedAt time.Time `json:"created_at"`
}
