package medicalrecord

import (
	"context"
	"v2/internal/domain/medicalrecord"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MedicalRecordPostgresRepository struct {
	db *pgxpool.Pool
}

func NewMedicalRecordPostgresRepository(db *pgxpool.Pool) *MedicalRecordPostgresRepository {
	return &MedicalRecordPostgresRepository{db: db}
}

func (r *MedicalRecordPostgresRepository) Create(ctx context.Context, mr *medicalrecord.MedicalRecord) error {
	if mr.ID == uuid.Nil {
		mr.ID = uuid.New()
	}
	_, err := r.db.Exec(ctx, `INSERT INTO medical_records (id, patient_id, mr_number, created_at) VALUES ($1, $2, $3, $4)`, mr.ID, mr.PatientID, mr.MRNumber, mr.CreatedAt)
	return err
}

func (r *MedicalRecordPostgresRepository) FindByPatientID(ctx context.Context, patientID uuid.UUID) (*medicalrecord.MedicalRecord, error) {
	row := r.db.QueryRow(ctx, `SELECT id, patient_id, mr_number, created_at FROM medical_records WHERE patient_id=$1`, patientID)
	var mr medicalrecord.MedicalRecord
	if err := row.Scan(&mr.ID, &mr.PatientID, &mr.MRNumber, &mr.CreatedAt); err != nil {
		return nil, err
	}
	return &mr, nil
}
