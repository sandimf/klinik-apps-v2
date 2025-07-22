package screening

import (
	"context"
	"v2/internal/domain/screening"

	"github.com/google/uuid"
)

type QuestionRepository interface {
	FindAll(ctx context.Context) ([]screening.ScreeningQuestion, error)
	Create(ctx context.Context, question *screening.ScreeningQuestion) error
	Update(ctx context.Context, id uuid.UUID, update map[string]interface{}) error
	FindByID(ctx context.Context, id uuid.UUID) (*screening.ScreeningQuestion, error)
}
