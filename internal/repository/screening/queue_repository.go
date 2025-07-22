package screening

import (
	"context"
	"v2/internal/domain/screening"

	"github.com/google/uuid"
)

type QueueRepository interface {
	Create(ctx context.Context, queue *screening.ScreeningQueue) error
	Update(ctx context.Context, id uuid.UUID, update map[string]interface{}) error
	FindAll(ctx context.Context) ([]screening.ScreeningQueue, error)
	FindByStatus(ctx context.Context, status string) ([]screening.ScreeningQueue, error)
	FindPaginatedByStatus(ctx context.Context, status string, page, limit int) ([]screening.ScreeningQueue, int64, error)
}
