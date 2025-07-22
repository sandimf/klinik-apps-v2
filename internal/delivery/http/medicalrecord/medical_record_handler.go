package medicalrecord

import (
	"v2/internal/usecase/medicalrecord"

	"github.com/gofiber/fiber/v2"
)

type MedicalRecordHandler struct {
	Usecase medicalrecord.MedicalRecordUsecase
}

func NewMedicalRecordHandler(u medicalrecord.MedicalRecordUsecase) *MedicalRecordHandler {
	return &MedicalRecordHandler{Usecase: u}
}

func (h *MedicalRecordHandler) CreateMedicalRecord(c *fiber.Ctx) error {
	var req struct {
		PatientID string `json:"patient_id"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse request"})
	}
	mr, err := h.Usecase.CreateMedicalRecord(c.Context(), req.PatientID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(mr)
}
