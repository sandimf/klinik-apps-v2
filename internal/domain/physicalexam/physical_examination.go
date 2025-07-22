package physicalexam

import (
	"time"

	"github.com/google/uuid"
)

type PhysicalExamination struct {
	ID                     uuid.UUID  `json:"id"`
	PatientID              uuid.UUID  `json:"patient_id"`
	ParamedisID            *uuid.UUID `json:"paramedis_id,omitempty"`
	DoctorID               *uuid.UUID `json:"doctor_id,omitempty"`
	BloodPressure          string     `json:"blood_pressure,omitempty"`
	HeartRate              *int       `json:"heart_rate,omitempty"`
	OxygenSaturation       *int       `json:"oxygen_saturation,omitempty"`
	RespiratoryRate        *int       `json:"respiratory_rate,omitempty"`
	BodyTemperature        *float64   `json:"body_temperature,omitempty"`
	PhysicalAssessment     string     `json:"physical_assessment,omitempty"`
	Reason                 string     `json:"reason,omitempty"`
	MedicalAdvice          string     `json:"medical_advice,omitempty"`
	HealthStatus           string     `json:"health_status,omitempty"`
	Pendampingan           []string   `json:"pendampingan,omitempty"`
	KonsultasiDokter       bool       `json:"konsultasi_dokter,omitempty"`
	KonsultasiDokterStatus string     `json:"konsultasi_dokter_status,omitempty"`
	DoctorAdvice           string     `json:"doctor_advice,omitempty"`
	CreatedAt              time.Time  `json:"created_at"`
	UpdatedAt              time.Time  `json:"updated_at"`
}
