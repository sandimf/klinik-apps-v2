package screening

import (
	"time"

	"github.com/google/uuid"
)

type ScreeningQueue struct {
	ID                uuid.UUID   `json:"id"`
	PatientInfo       PatientInfo `json:"patient_info"`
	ScreeningAnswerID uuid.UUID   `json:"screening_answer_id"`
	Status            string      `json:"status"` // waiting, in_progress, done
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
}
