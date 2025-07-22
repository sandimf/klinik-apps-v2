package screening

import (
	"context"
	"encoding/json"
	"v2/internal/domain/screening"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type QueuePostgresRepository struct {
	db *pgxpool.Pool
}

func NewQueuePostgresRepository(db *pgxpool.Pool) *QueuePostgresRepository {
	return &QueuePostgresRepository{db: db}
}

func (r *QueuePostgresRepository) Create(ctx context.Context, q *screening.ScreeningQueue) error {
	if q.ID == uuid.Nil {
		q.ID = uuid.New()
	}
	patientInfo, _ := json.Marshal(q.PatientInfo)
	_, err := r.db.Exec(ctx, `INSERT INTO screening_queues (id, patient_info, screening_answer_id, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`, q.ID, patientInfo, q.ScreeningAnswerID, q.Status, q.CreatedAt, q.UpdatedAt)
	return err
}

func (r *QueuePostgresRepository) Update(ctx context.Context, id uuid.UUID, update map[string]interface{}) error {
	// Sederhana: hanya update status
	_, err := r.db.Exec(ctx, `UPDATE screening_queues SET status=$1, updated_at=NOW() WHERE id=$2`, update["status"], id)
	return err
}

func (r *QueuePostgresRepository) FindAll(ctx context.Context) ([]screening.ScreeningQueue, error) {
	rows, err := r.db.Query(ctx, `SELECT id, patient_info, screening_answer_id, status, created_at, updated_at FROM screening_queues`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []screening.ScreeningQueue
	for rows.Next() {
		var q screening.ScreeningQueue
		var patientInfoData []byte
		if err := rows.Scan(&q.ID, &patientInfoData, &q.ScreeningAnswerID, &q.Status, &q.CreatedAt, &q.UpdatedAt); err != nil {
			return nil, err
		}
		_ = json.Unmarshal(patientInfoData, &q.PatientInfo)
		result = append(result, q)
	}
	return result, nil
}

func (r *QueuePostgresRepository) FindByStatus(ctx context.Context, status string) ([]screening.ScreeningQueue, error) {
	rows, err := r.db.Query(ctx, `SELECT id, patient_info, screening_answer_id, status, created_at, updated_at FROM screening_queues WHERE status=$1`, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []screening.ScreeningQueue
	for rows.Next() {
		var q screening.ScreeningQueue
		var patientInfoData []byte
		if err := rows.Scan(&q.ID, &patientInfoData, &q.ScreeningAnswerID, &q.Status, &q.CreatedAt, &q.UpdatedAt); err != nil {
			return nil, err
		}
		_ = json.Unmarshal(patientInfoData, &q.PatientInfo)
		result = append(result, q)
	}
	return result, nil
}

func (r *QueuePostgresRepository) FindPaginatedByStatus(ctx context.Context, status string, page, limit int) ([]screening.ScreeningQueue, int64, error) {
	offset := (page - 1) * limit
	rows, err := r.db.Query(ctx, `SELECT id, patient_info, screening_answer_id, status, created_at, updated_at FROM screening_queues WHERE status=$1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`, status, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	var result []screening.ScreeningQueue
	for rows.Next() {
		var q screening.ScreeningQueue
		var patientInfoData []byte
		if err := rows.Scan(&q.ID, &patientInfoData, &q.ScreeningAnswerID, &q.Status, &q.CreatedAt, &q.UpdatedAt); err != nil {
			return nil, 0, err
		}
		_ = json.Unmarshal(patientInfoData, &q.PatientInfo)
		result = append(result, q)
	}
	// Hitung total
	var total int64
	row := r.db.QueryRow(ctx, `SELECT COUNT(*) FROM screening_queues WHERE status=$1`, status)
	if err := row.Scan(&total); err != nil {
		return nil, 0, err
	}
	return result, total, nil
}
