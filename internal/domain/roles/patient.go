package roles

import (
	"time"

	"github.com/google/uuid"
)

type Patient struct {
	ID          uuid.UUID `json:"id"`
	NIK         string    `json:"nik"`
	FullName    string    `json:"full_name"`
	BirthPlace  string    `json:"birth_place"`
	BirthDate   string    `json:"birth_date"`
	Gender      string    `json:"gender"`
	Address     string    `json:"address"`
	RT          string    `json:"rt"`
	RW          string    `json:"rw"`
	Village     string    `json:"village"`
	District    string    `json:"district"`
	Religion    string    `json:"religion"`
	Marital     string    `json:"marital"`
	Job         string    `json:"job"`
	Nationality string    `json:"nationality"`
	ValidUntil  string    `json:"valid_until"`
	BloodType   string    `json:"blood_type"`
	Height      int       `json:"height"`
	Weight      int       `json:"weight"`
	Age         int       `json:"age"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	KTPImages   []string  `json:"ktp_images,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
