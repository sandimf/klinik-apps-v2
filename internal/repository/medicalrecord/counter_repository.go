package medicalrecord

import "context"

type CounterRepository interface {
	GetNextSequence(ctx context.Context, key string) (int64, error)
}
