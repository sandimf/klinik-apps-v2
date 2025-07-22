package screening

import (
	"context"
	"encoding/json"
	"v2/internal/domain/screening"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AnswerPostgresRepository struct {
	db *pgxpool.Pool
}

func NewAnswerPostgresRepository(db *pgxpool.Pool) *AnswerPostgresRepository {
	return &AnswerPostgresRepository{db: db}
}

func (r *AnswerPostgresRepository) Create(ctx context.Context, a *screening.ScreeningAnswer) error {
	if a.ID == uuid.Nil {
		a.ID = uuid.New()
	}
	patientInfo, _ := json.Marshal(a.PatientInfo)
	answers, _ := json.Marshal(a.Answers)
	_, err := r.db.Exec(ctx, `INSERT INTO screening_answers (id, patient_info, answers, created_at) VALUES ($1, $2, $3, $4)`, a.ID, patientInfo, answers, a.CreatedAt)
	return err
}

func (r *AnswerPostgresRepository) Update(ctx context.Context, id uuid.UUID, update map[string]interface{}) error {
	// Sederhana: hanya update answers
	answers, _ := json.Marshal(update["answers"])
	_, err := r.db.Exec(ctx, `UPDATE screening_answers SET answers=$1 WHERE id=$2`, answers, id)
	return err
}

func (r *AnswerPostgresRepository) FindByID(ctx context.Context, id uuid.UUID) (*screening.ScreeningAnswer, error) {
	row := r.db.QueryRow(ctx, `SELECT id, patient_info, answers, created_at FROM screening_answers WHERE id=$1`, id)
	var a screening.ScreeningAnswer
	var patientInfoData, answersData []byte
	if err := row.Scan(&a.ID, &patientInfoData, &answersData, &a.CreatedAt); err != nil {
		return nil, err
	}
	_ = json.Unmarshal(patientInfoData, &a.PatientInfo)
	_ = json.Unmarshal(answersData, &a.Answers)
	return &a, nil
}
