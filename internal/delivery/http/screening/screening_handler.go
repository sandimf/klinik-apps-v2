package screening

import (
	"v2/internal/domain/screening"
	usecase "v2/internal/usecase/screening"

	"math"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ScreeningHandler struct {
	Usecase usecase.ScreeningUsecase
}

func NewScreeningHandler(u usecase.ScreeningUsecase) *ScreeningHandler {
	return &ScreeningHandler{Usecase: u}
}

func (h *ScreeningHandler) GetQuestions(c *fiber.Ctx) error {
	questions, err := h.Usecase.GetQuestions(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to get questions"})
	}
	return c.JSON(questions)
}

func (h *ScreeningHandler) SubmitAnswer(c *fiber.Ctx) error {
	var req screening.ScreeningAnswer
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse request"})
	}
	if err := h.Usecase.SubmitAnswer(c.Context(), &req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to submit answer"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "answer submitted"})
}

func (h *ScreeningHandler) EnqueueScreening(c *fiber.Ctx) error {
	var req screening.ScreeningQueue
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse request"})
	}
	if err := h.Usecase.EnqueueScreening(c.Context(), &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "enqueued"})
}

func (h *ScreeningHandler) ScreeningWithPatient(c *fiber.Ctx) error {
	// Parsing multipart/form-data (data pasien + screening + file KTP opsional)
	var req struct {
		Patient   map[string]interface{} `json:"patient" form:"patient"`
		Screening map[string]interface{} `json:"screening" form:"screening"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse request"})
	}
	// TODO: mapping req.Patient dan req.Screening ke struct
	// TODO: handle file upload jika ada
	// TODO: panggil usecase ScreeningWithPatient
	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{"message": "not implemented"})
}

func (h *ScreeningHandler) UpdateScreeningAnswer(c *fiber.Ctx) error {
	id := c.Params("id")
	var update map[string]interface{}
	if err := c.BodyParser(&update); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse request"})
	}
	if err := h.Usecase.UpdateScreeningAnswer(c.Context(), id, update); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "screening answer updated"})
}

func (h *ScreeningHandler) CreateQuestion(c *fiber.Ctx) error {
	var q screening.ScreeningQuestion
	if err := c.BodyParser(&q); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse request"})
	}
	if err := h.Usecase.CreateQuestion(c.Context(), &q); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(q)
}

func (h *ScreeningHandler) UpdateQuestion(c *fiber.Ctx) error {
	id := c.Params("id")
	var update map[string]interface{}
	if err := c.BodyParser(&update); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse request"})
	}
	if err := h.Usecase.UpdateQuestion(c.Context(), id, update); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "question updated"})
}

func (h *ScreeningHandler) ListQueue(c *fiber.Ctx) error {
	status := c.Query("status", "screening_pending")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	queues, total, err := h.Usecase.FindQueuePaginatedByStatus(c.Context(), status, page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	totalPages := int(math.Ceil(float64(total) / float64(limit)))
	return c.JSON(fiber.Map{
		"data": queues,
		"meta": fiber.Map{
			"page":        page,
			"limit":       limit,
			"total":       total,
			"total_pages": totalPages,
		},
	})
}
