package roles

import (
	"path/filepath"
	"time"
	"v2/internal/domain/roles"
	usecase "v2/internal/usecase/roles"

	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PatientHandler struct {
	Usecase usecase.PatientUsecase
}

func NewPatientHandler(u usecase.PatientUsecase) *PatientHandler {
	return &PatientHandler{Usecase: u}
}

func (h *PatientHandler) CreateOrUpdatePatient(c *fiber.Ctx) error {
	patient := new(roles.Patient)
	if err := c.BodyParser(patient); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse request"})
	}

	// Handle file upload (ktp_images)
	form, err := c.MultipartForm()
	if err == nil && form.File != nil {
		files := form.File["ktp_images"]
		var paths []string
		for _, file := range files {
			filename := time.Now().Format("20060102150405") + "_" + file.Filename
			path := filepath.Join("public", "uploads", filename)
			if err := c.SaveFile(file, path); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to save file"})
			}
			paths = append(paths, "/"+path)
		}
		patient.KTPImages = paths
	}

	updated, err := h.Usecase.CreateOrUpdatePatient(c.Context(), patient)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(updated)
}

func (h *PatientHandler) GetAll(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	patients, total, err := h.Usecase.FindAllPaginated(c.Context(), page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	totalPages := int(math.Ceil(float64(total) / float64(limit)))
	return c.JSON(fiber.Map{
		"data": patients,
		"meta": fiber.Map{
			"page":        page,
			"limit":       limit,
			"total":       total,
			"total_pages": totalPages,
		},
	})
}
