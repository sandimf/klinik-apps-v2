package repository

import (
	"context"
	"v2/internal/domain"
)

type CashierRepository interface {
	Create(ctx context.Context, cashier *domain.Cashier) error
}
