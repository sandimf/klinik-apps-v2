package roles

import (
	"v2/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserRepo repository.UserRepository
}

func NewUserHandler(r repository.UserRepository) *UserHandler {
	return &UserHandler{UserRepo: r}
}

func (h *UserHandler) Me(c *fiber.Ctx) error {
	userID := c.Locals("user_id")
	if userID == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}
	user, err := h.UserRepo.FindByID(c.Context(), userID.(string))
	if err != nil || user == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	}
	return c.JSON(user)
}
