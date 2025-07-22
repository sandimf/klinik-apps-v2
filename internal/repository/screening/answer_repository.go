package screening

import (
	"context"
	"v2/internal/domain/screening"

	"github.com/google/uuid"
)

type AnswerRepository interface {
	Create(ctx context.Context, answer *screening.ScreeningAnswer) error
	Update(ctx context.Context, id uuid.UUID, update map[string]interface{}) error
}
