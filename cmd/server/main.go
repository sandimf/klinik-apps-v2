package main

import (
	"log"
	"v2/internal/config"
	"v2/internal/delivery/http"
	screeningHandlerPkg "v2/internal/delivery/http/screening"
	"v2/internal/repository"
	patientRepoPkg "v2/internal/repository/roles"
	"v2/internal/repository/screening"
	"v2/internal/usecase"
	screeningUsecasePkg "v2/internal/usecase/screening"
	"v2/internal/utils"

	_ "v2/docs" // ganti dengan module path Anda jika berbeda

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func main() {
	// 1. Load Config
	cfg := config.LoadConfig()

	// 2. Connect to PostgreSQL
	pgPool, err := utils.NewPostgresDB()
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer pgPool.Close()

	// 3. Dependency Injection
	userRepo := repository.NewUserPostgresRepository(pgPool)
	patientRepo := patientRepoPkg.NewPatientPostgresRepository(pgPool)
	userUsecase := usecase.NewUserUsecase(userRepo, patientRepo)
	userHandler := http.NewUserHandler(userUsecase, userRepo)

	// Screening
	questionRepo := screening.NewQuestionPostgresRepository(pgPool)
	answerRepo := screening.NewAnswerPostgresRepository(pgPool)
	queueRepo := screening.NewQueuePostgresRepository(pgPool)
	screeningUsecase := screeningUsecasePkg.NewScreeningUsecase(questionRepo, answerRepo, queueRepo, patientRepo, userRepo)
	screeningHandler := screeningHandlerPkg.NewScreeningHandler(screeningUsecase)

	// TODO: Ganti semua repository dan usecase lain ke versi Postgres jika sudah ada
	// Sementara, screening, medical record, dsb masih pakai Mongo jika belum dimigrasi

	// 4. Setup Fiber & Register Routes
	app := fiber.New()
	app.Use(cors.New())

	// Health check endpoint
	app.Get("/ping", func(c *fiber.Ctx) error {
		log.Println("Ping endpoint hit")
		return c.SendString("pong")
	})

	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	api := app.Group("/api/v1")
	http.RegisterRoutes(api, userHandler, screeningHandler, nil, nil, nil, nil) // TODO: inject handler lain jika sudah migrasi

	// 5. Start Server
	port := cfg.Port
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}
