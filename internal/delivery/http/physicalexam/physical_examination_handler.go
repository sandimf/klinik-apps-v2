package physicalexam

import (
	"v2/internal/domain/physicalexam"
	usecase "v2/internal/usecase/physicalexam"

	"github.com/gofiber/fiber/v2"
)

type PhysicalExaminationHandler struct {
	Usecase usecase.PhysicalExaminationUsecase
}

func NewPhysicalExaminationHandler(u usecase.PhysicalExaminationUsecase) *PhysicalExaminationHandler {
	return &PhysicalExaminationHandler{Usecase: u}
}

func (h *PhysicalExaminationHandler) Create(c *fiber.Ctx) error {
	exam := new(physicalexam.PhysicalExamination)
	if err := c.BodyParser(exam); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse request"})
	}
	if err := h.Usecase.Create(c.Context(), exam); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(exam)
}

func (h *PhysicalExaminationHandler) GetByPatientID(c *fiber.Ctx) error {
	patientID := c.Query("patient_id")
	if patientID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "patient_id required"})
	}
	exams, err := h.Usecase.FindByPatientID(c.Context(), patientID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(exams)
}

func (h *PhysicalExaminationHandler) GetDoctorConsultations(c *fiber.Ctx) error {
	exams, err := h.Usecase.FindDoctorConsultations(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(exams)
}

func (h *PhysicalExaminationHandler) UpdateConsultationStatus(c *fiber.Ctx) error {
	id := c.Params("id")
	var req struct {
		Status string `json:"status"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse request"})
	}
	if err := h.Usecase.UpdateConsultationStatus(c.Context(), id, req.Status); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "consultation status updated"})
}

func (h *PhysicalExaminationHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var update map[string]interface{}
	if err := c.BodyParser(&update); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse request"})
	}
	if err := h.Usecase.Update(c.Context(), id, update); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "physical examination updated"})
}
