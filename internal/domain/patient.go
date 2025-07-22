package domain

import "github.com/google/uuid"

type Patient struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	FullName    string    `json:"full_name"`
	NIK         string    `json:"nik"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
}
