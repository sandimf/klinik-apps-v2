package http

import (
	"log"
	"strings"
	"v2/internal/usecase"

	"v2/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserUsecase usecase.UserUsecase
	UserRepo    repository.UserRepository
}

func NewUserHandler(u usecase.UserUsecase, r repository.UserRepository) *UserHandler {
	return &UserHandler{UserUsecase: u, UserRepo: r}
}

// DTO for patient registration request
type RegisterPatientRequest struct {
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=8"`
	FullName    string `json:"full_name" validate:"required"`
	NIK         string `json:"nik"`
	Phone       string `json:"phone"`
	BirthPlace  string `json:"birth_place"`
	BirthDate   string `json:"birth_date"`
	Gender      string `json:"gender"`
	Address     string `json:"address"`
	RT          string `json:"rt"`
	RW          string `json:"rw"`
	Village     string `json:"village"`
	District    string `json:"district"`
	Religion    string `json:"religion"`
	Marital     string `json:"marital"`
	Job         string `json:"job"`
	Nationality string `json:"nationality"`
	ValidUntil  string `json:"valid_until"`
	BloodType   string `json:"blood_type"`
	Height      int    `json:"height"`
	Weight      int    `json:"weight"`
	Age         int    `json:"age"`
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	log.Println("Masuk handler /register")
	var req RegisterPatientRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse request"})
	}

	input := usecase.RegisterPatientInput{
		Email:       req.Email,
		Password:    req.Password,
		FullName:    req.FullName,
		NIK:         req.NIK,
		Phone:       req.Phone,
		BirthPlace:  req.BirthPlace,
		BirthDate:   req.BirthDate,
		Gender:      req.Gender,
		Address:     req.Address,
		RT:          req.RT,
		RW:          req.RW,
		Village:     req.Village,
		District:    req.District,
		Religion:    req.Religion,
		Marital:     req.Marital,
		Job:         req.Job,
		Nationality: req.Nationality,
		ValidUntil:  req.ValidUntil,
		BloodType:   req.BloodType,
		Height:      req.Height,
		Weight:      req.Weight,
		Age:         req.Age,
	}

	err := h.UserUsecase.RegisterPatient(c.Context(), input)
	if err != nil {
		if strings.Contains(err.Error(), "email already exists") {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": "email already exists"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to register patient"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "patient registered successfully"})
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	log.Println("Masuk handler /login")
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse request"})
	}
	// TODO: Add validation for the request struct
	token, err := h.UserUsecase.Login(c.Context(), req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid email or password"})
	}
	return c.JSON(LoginResponse{Token: token})
}

// Me godoc
// @Summary Get current user profile
// @Description Get data of the currently logged-in user
// @Tags Profile
// @Accept json
// @Produce json
// @Success 200 {object} roles.User
// @Router /api/v1/me [get]
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
