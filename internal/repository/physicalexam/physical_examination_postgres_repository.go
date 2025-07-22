package physicalexam

import (
	"context"
	"v2/internal/domain/physicalexam"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PhysicalExaminationPostgresRepository struct {
	db *pgxpool.Pool
}

func NewPhysicalExaminationPostgresRepository(db *pgxpool.Pool) *PhysicalExaminationPostgresRepository {
	return &PhysicalExaminationPostgresRepository{db: db}
}

func (r *PhysicalExaminationPostgresRepository) Create(ctx context.Context, exam *physicalexam.PhysicalExamination) error {
	if exam.ID == uuid.Nil {
		exam.ID = uuid.New()
	}
	_, err := r.db.Exec(ctx, `INSERT INTO physical_examinations (id, patient_id, paramedis_id, doctor_id, blood_pressure, heart_rate, oxygen_saturation, respiratory_rate, body_temperature, physical_assessment, reason, medical_advice, health_status, pendampingan, konsultasi_dokter, konsultasi_dokter_status, doctor_advice, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19)`, exam.ID, exam.PatientID, exam.ParamedisID, exam.DoctorID, exam.BloodPressure, exam.HeartRate, exam.OxygenSaturation, exam.RespiratoryRate, exam.BodyTemperature, exam.PhysicalAssessment, exam.Reason, exam.MedicalAdvice, exam.HealthStatus, exam.Pendampingan, exam.KonsultasiDokter, exam.KonsultasiDokterStatus, exam.DoctorAdvice, exam.CreatedAt, exam.UpdatedAt)
	return err
}

func (r *PhysicalExaminationPostgresRepository) FindByPatientID(ctx context.Context, patientID uuid.UUID) ([]physicalexam.PhysicalExamination, error) {
	rows, err := r.db.Query(ctx, `SELECT id, patient_id, paramedis_id, doctor_id, blood_pressure, heart_rate, oxygen_saturation, respiratory_rate, body_temperature, physical_assessment, reason, medical_advice, health_status, pendampingan, konsultasi_dokter, konsultasi_dokter_status, doctor_advice, created_at, updated_at FROM physical_examinations WHERE patient_id=$1`, patientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []physicalexam.PhysicalExamination
	for rows.Next() {
		var e physicalexam.PhysicalExamination
		if err := rows.Scan(&e.ID, &e.PatientID, &e.ParamedisID, &e.DoctorID, &e.BloodPressure, &e.HeartRate, &e.OxygenSaturation, &e.RespiratoryRate, &e.BodyTemperature, &e.PhysicalAssessment, &e.Reason, &e.MedicalAdvice, &e.HealthStatus, &e.Pendampingan, &e.KonsultasiDokter, &e.KonsultasiDokterStatus, &e.DoctorAdvice, &e.CreatedAt, &e.UpdatedAt); err != nil {
			return nil, err
		}
		result = append(result, e)
	}
	return result, nil
}

func (r *PhysicalExaminationPostgresRepository) FindDoctorConsultations(ctx context.Context) ([]physicalexam.PhysicalExamination, error) {
	rows, err := r.db.Query(ctx, `SELECT id, patient_id, paramedis_id, doctor_id, blood_pressure, heart_rate, oxygen_saturation, respiratory_rate, body_temperature, physical_assessment, reason, medical_advice, health_status, pendampingan, konsultasi_dokter, konsultasi_dokter_status, doctor_advice, created_at, updated_at FROM physical_examinations WHERE konsultasi_dokter=true`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []physicalexam.PhysicalExamination
	for rows.Next() {
		var e physicalexam.PhysicalExamination
		if err := rows.Scan(&e.ID, &e.PatientID, &e.ParamedisID, &e.DoctorID, &e.BloodPressure, &e.HeartRate, &e.OxygenSaturation, &e.RespiratoryRate, &e.BodyTemperature, &e.PhysicalAssessment, &e.Reason, &e.MedicalAdvice, &e.HealthStatus, &e.Pendampingan, &e.KonsultasiDokter, &e.KonsultasiDokterStatus, &e.DoctorAdvice, &e.CreatedAt, &e.UpdatedAt); err != nil {
			return nil, err
		}
		result = append(result, e)
	}
	return result, nil
}

func (r *PhysicalExaminationPostgresRepository) UpdateConsultationStatus(ctx context.Context, id uuid.UUID, status string) error {
	_, err := r.db.Exec(ctx, `UPDATE physical_examinations SET konsultasi_dokter_status=$1 WHERE id=$2`, status, id)
	return err
}

func (r *PhysicalExaminationPostgresRepository) Update(ctx context.Context, id uuid.UUID, update map[string]interface{}) error {
	// Sederhana: hanya update beberapa field
	_, err := r.db.Exec(ctx, `UPDATE physical_examinations SET blood_pressure=$1, heart_rate=$2, oxygen_saturation=$3, respiratory_rate=$4, body_temperature=$5, physical_assessment=$6, reason=$7, medical_advice=$8, health_status=$9, pendampingan=$10, konsultasi_dokter=$11, konsultasi_dokter_status=$12, doctor_advice=$13, updated_at=NOW() WHERE id=$14`, update["blood_pressure"], update["heart_rate"], update["oxygen_saturation"], update["respiratory_rate"], update["body_temperature"], update["physical_assessment"], update["reason"], update["medical_advice"], update["health_status"], update["pendampingan"], update["konsultasi_dokter"], update["konsultasi_dokter_status"], update["doctor_advice"], id)
	return err
}
