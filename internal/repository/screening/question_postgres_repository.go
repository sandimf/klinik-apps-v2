package screening

import (
	"context"
	"v2/internal/domain/screening"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type QuestionPostgresRepository struct {
	db *pgxpool.Pool
}

func NewQuestionPostgresRepository(db *pgxpool.Pool) *QuestionPostgresRepository {
	return &QuestionPostgresRepository{db: db}
}

func (r *QuestionPostgresRepository) FindAll(ctx context.Context) ([]screening.ScreeningQuestion, error) {
	rows, err := r.db.Query(ctx, `SELECT id, label, type, options FROM screening_questions`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []screening.ScreeningQuestion
	for rows.Next() {
		var q screening.ScreeningQuestion
		var options []string
		if err := rows.Scan(&q.ID, &q.Label, &q.Type, &options); err != nil {
			return nil, err
		}
		q.Options = options
		result = append(result, q)
	}
	return result, nil
}

func (r *QuestionPostgresRepository) Create(ctx context.Context, q *screening.ScreeningQuestion) error {
	if q.ID == uuid.Nil {
		q.ID = uuid.New()
	}
	_, err := r.db.Exec(ctx, `INSERT INTO screening_questions (id, label, type, options) VALUES ($1, $2, $3, $4)`, q.ID, q.Label, q.Type, q.Options)
	return err
}

func (r *QuestionPostgresRepository) Update(ctx context.Context, id uuid.UUID, update map[string]interface{}) error {
	// Sederhana: hanya update label, type, options
	_, err := r.db.Exec(ctx, `UPDATE screening_questions SET label=$1, type=$2, options=$3 WHERE id=$4`, update["label"], update["type"], update["options"], id)
	return err
}

func (r *QuestionPostgresRepository) FindByID(ctx context.Context, id uuid.UUID) (*screening.ScreeningQuestion, error) {
	row := r.db.QueryRow(ctx, `SELECT id, label, type, options FROM screening_questions WHERE id=$1`, id)
	var q screening.ScreeningQuestion
	var options []string
	if err := row.Scan(&q.ID, &q.Label, &q.Type, &options); err != nil {
		return nil, err
	}
	q.Options = options
	return &q, nil
}
