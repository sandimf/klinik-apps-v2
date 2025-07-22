package medicine

import (
	"time"

	"github.com/google/uuid"
)

type Medicine struct {
	ID           uuid.UUID `json:"id"`
	Barcode      string    `json:"barcode"`
	MedicineName string    `json:"medicine_name"`
	BrandName    string    `json:"brand_name"`
	Category     string    `json:"category"`
	Dosage       int       `json:"dosage"`
	Content      string    `json:"content"`
	Quantity     int       `json:"quantity"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
