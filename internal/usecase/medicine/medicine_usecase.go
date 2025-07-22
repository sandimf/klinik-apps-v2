package medicine

import (
	"context"
	"v2/internal/domain/medicine"
	repo "v2/internal/repository/medicine"
)

type MedicineUsecase interface {
	Create(ctx context.Context, medicine *medicine.Medicine) error
	Update(ctx context.Context, id string, update map[string]interface{}) error
	FindAll(ctx context.Context) ([]medicine.Medicine, error)
	FindAllPaginated(ctx context.Context, page, limit int) ([]medicine.Medicine, int64, error)
}

type medicineUsecase struct {
	repo repo.MedicineRepository
}

func NewMedicineUsecase(r repo.MedicineRepository) MedicineUsecase {
	return &medicineUsecase{repo: r}
}

func (u *medicineUsecase) Create(ctx context.Context, medicine *medicine.Medicine) error {
	return u.repo.Create(ctx, medicine)
}

func (u *medicineUsecase) Update(ctx context.Context, id string, update map[string]interface{}) error {
	return u.repo.Update(ctx, id, update)
}

func (u *medicineUsecase) FindAll(ctx context.Context) ([]medicine.Medicine, error) {
	return u.repo.FindAll(ctx)
}

func (u *medicineUsecase) FindAllPaginated(ctx context.Context, page, limit int) ([]medicine.Medicine, int64, error) {
	return u.repo.FindAllPaginated(ctx, page, limit)
}
