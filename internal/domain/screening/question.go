package screening

import "github.com/google/uuid"

type ScreeningQuestion struct {
	ID      uuid.UUID `json:"id"`
	Label   string    `json:"label"`
	Type    string    `json:"type"`
	Options []string  `json:"options,omitempty"`
}
