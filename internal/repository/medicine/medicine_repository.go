package medicine

import (
	"context"
	"v2/internal/domain/medicine"
)

type MedicineRepository interface {
	Create(ctx context.Context, medicine *medicine.Medicine) error
	Update(ctx context.Context, id string, update map[string]interface{}) error
	FindAll(ctx context.Context) ([]medicine.Medicine, error)
	FindAllPaginated(ctx context.Context, page, limit int) ([]medicine.Medicine, int64, error)
}
