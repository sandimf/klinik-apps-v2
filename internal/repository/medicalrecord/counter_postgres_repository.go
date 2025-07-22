package medicalrecord

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CounterPostgresRepository struct {
	db *pgxpool.Pool
}

func NewCounterPostgresRepository(db *pgxpool.Pool) *CounterPostgresRepository {
	return &CounterPostgresRepository{db: db}
}

func (r *CounterPostgresRepository) GetNextSequence(ctx context.Context, name string) (int, error) {
	// Upsert counter row, increment seq, return new value
	_, err := r.db.Exec(ctx, `INSERT INTO counters (name, seq) VALUES ($1, 1) ON CONFLICT (name) DO UPDATE SET seq = counters.seq + 1`, name)
	if err != nil {
		return 0, err
	}
	var seq int
	row := r.db.QueryRow(ctx, `SELECT seq FROM counters WHERE name=$1`, name)
	if err := row.Scan(&seq); err != nil {
		return 0, err
	}
	return seq, nil
}
