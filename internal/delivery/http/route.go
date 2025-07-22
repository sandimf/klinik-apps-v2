package http

import (
	medicalRecordHandlerPkg "v2/internal/delivery/http/medicalrecord"
	medicineHandlerPkg "v2/internal/delivery/http/medicine"
	physicalExamHandlerPkg "v2/internal/delivery/http/physicalexam"
	patientHandlerPkg "v2/internal/delivery/http/roles"
	screeningHandlerPkg "v2/internal/delivery/http/screening"
	"v2/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(router fiber.Router, userHandler *UserHandler, screeningHandler *screeningHandlerPkg.ScreeningHandler, medicalRecordHandler *medicalRecordHandlerPkg.MedicalRecordHandler, patientHandler *patientHandlerPkg.PatientHandler, physicalExamHandler *physicalExamHandlerPkg.PhysicalExaminationHandler, medicineHandler *medicineHandlerPkg.MedicineHandler) {
	router.Post("/register", userHandler.Register)
	router.Post("/login", userHandler.Login)
	router.Get("/me", middleware.AuthMiddleware(), userHandler.Me)

	// Patient
	router.Post("/patients", patientHandler.CreateOrUpdatePatient)

	// Screening routes (no auth)
	router.Get("/screening/questions", screeningHandler.GetQuestions)
	router.Post("/screening/questions", middleware.AuthMiddleware(), middleware.AdminOnly(), screeningHandler.CreateQuestion)
	router.Patch("/screening/questions/:id", middleware.AuthMiddleware(), middleware.AdminOnly(), screeningHandler.UpdateQuestion)
	router.Post("/screening/answers", screeningHandler.SubmitAnswer)
	router.Post("/screening/queue", screeningHandler.EnqueueScreening)
	router.Post("/screening/with-patient", screeningHandler.ScreeningWithPatient)
	router.Get("/screening/queue", screeningHandler.ListQueue)

	// Medical Record
	router.Post("/medical-record", medicalRecordHandler.CreateMedicalRecord)

	// Physical Examination
	router.Post("/physical-examinations", physicalExamHandler.Create)
	router.Get("/physical-examinations/by-patient", physicalExamHandler.GetByPatientID)
	router.Get("/doctor/consultations", physicalExamHandler.GetDoctorConsultations)
	router.Patch("/physical-examinations/:id/consultation-status", physicalExamHandler.UpdateConsultationStatus)
	router.Patch("/physical-examinations/:id", physicalExamHandler.Update)
	router.Patch("/screening/answers/:id", screeningHandler.UpdateScreeningAnswer)
	router.Get("/doctor/patients", patientHandler.GetAll)

	// Medicine
	router.Post("/medicines", middleware.AuthMiddleware(), middleware.AdminOnly(), medicineHandler.Create)
	router.Patch("/medicines/:id", middleware.AuthMiddleware(), middleware.AdminOnly(), medicineHandler.Update)
	router.Get("/medicines", medicineHandler.FindAll)
}
