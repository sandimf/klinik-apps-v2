package medicine

import (
	"v2/internal/domain/medicine"
	usecase "v2/internal/usecase/medicine"

	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type MedicineHandler struct {
	Usecase usecase.MedicineUsecase
}

func NewMedicineHandler(u usecase.MedicineUsecase) *MedicineHandler {
	return &MedicineHandler{Usecase: u}
}

func (h *MedicineHandler) Create(c *fiber.Ctx) error {
	var m medicine.Medicine
	if err := c.BodyParser(&m); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse request"})
	}
	if err := h.Usecase.Create(c.Context(), &m); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(m)
}

func (h *MedicineHandler) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var update map[string]interface{}
	if err := c.BodyParser(&update); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse request"})
	}
	if err := h.Usecase.Update(c.Context(), id, update); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "medicine updated"})
}

// FindAll godoc
// @Summary List Obat
// @Description Mendapatkan daftar obat (pagination)
// @Tags Medicines
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Page size"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/medicines [get]
func (h *MedicineHandler) FindAll(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	medicines, total, err := h.Usecase.FindAllPaginated(c.Context(), page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	totalPages := int(math.Ceil(float64(total) / float64(limit)))
	return c.JSON(fiber.Map{
		"data": medicines,
		"meta": fiber.Map{
			"page":        page,
			"limit":       limit,
			"total":       total,
			"total_pages": totalPages,
		},
	})
}
