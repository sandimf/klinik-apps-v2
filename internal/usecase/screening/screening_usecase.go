package screening

import (
	"context"
	"errors"
	"time"
	rolesdomain "v2/internal/domain/roles"
	"v2/internal/domain/screening"
	userrepo "v2/internal/repository"
	rolesrepo "v2/internal/repository/roles"
	repo "v2/internal/repository/screening"
	"v2/internal/utils"

	"github.com/google/uuid"
)

type ScreeningUsecase interface {
	GetQuestions(ctx context.Context) ([]screening.ScreeningQuestion, error)
	SubmitAnswer(ctx context.Context, answer *screening.ScreeningAnswer) error
	EnqueueScreening(ctx context.Context, queue *screening.ScreeningQueue) error
	UpdateScreeningAnswer(ctx context.Context, id string, update map[string]interface{}) error
	CreateQuestion(ctx context.Context, question *screening.ScreeningQuestion) error
	UpdateQuestion(ctx context.Context, id string, update map[string]interface{}) error
	FindQueuePaginatedByStatus(ctx context.Context, status string, page, limit int) ([]screening.ScreeningQueue, int64, error)
}

type screeningUsecase struct {
	questionRepo repo.QuestionRepository
	answerRepo   repo.AnswerRepository
	queueRepo    repo.QueueRepository
	patientRepo  rolesrepo.PatientRepository
	userRepo     userrepo.UserRepository
}

func NewScreeningUsecase(qr repo.QuestionRepository, ar repo.AnswerRepository, qrq repo.QueueRepository, pr rolesrepo.PatientRepository, ur userrepo.UserRepository) ScreeningUsecase {
	return &screeningUsecase{
		questionRepo: qr,
		answerRepo:   ar,
		queueRepo:    qrq,
		patientRepo:  pr,
		userRepo:     ur,
	}
}

func (u *screeningUsecase) GetQuestions(ctx context.Context) ([]screening.ScreeningQuestion, error) {
	return u.questionRepo.FindAll(ctx)
}

func (u *screeningUsecase) SubmitAnswer(ctx context.Context, answer *screening.ScreeningAnswer) error {
	answer.CreatedAt = time.Now()
	return u.answerRepo.Create(ctx, answer)
}

func (u *screeningUsecase) EnqueueScreening(ctx context.Context, queue *screening.ScreeningQueue) error {
	if queue.ScreeningAnswerID == uuid.Nil {
		return errors.New("screening_answer_id required")
	}
	return u.queueRepo.Create(ctx, queue)
}

func (u *screeningUsecase) UpdateScreeningAnswer(ctx context.Context, id string, update map[string]interface{}) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return u.answerRepo.Update(ctx, uid, update)
}

func (u *screeningUsecase) CreateQuestion(ctx context.Context, question *screening.ScreeningQuestion) error {
	return u.questionRepo.Create(ctx, question)
}

func (u *screeningUsecase) UpdateQuestion(ctx context.Context, id string, update map[string]interface{}) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	return u.questionRepo.Update(ctx, uid, update)
}

func (u *screeningUsecase) FindQueuePaginatedByStatus(ctx context.Context, status string, page, limit int) ([]screening.ScreeningQueue, int64, error) {
	return u.queueRepo.FindPaginatedByStatus(ctx, status, page, limit)
}

type ScreeningWithPatientInput struct {
	Patient   *rolesdomain.Patient
	Screening screening.ScreeningAnswer
}

func (u *screeningUsecase) ScreeningWithPatient(ctx context.Context, input ScreeningWithPatientInput) error {
	// 1. Cek pasien by NIK/email
	pasien, _ := u.patientRepo.FindByNIK(ctx, input.Patient.NIK)
	if pasien == nil {
		// Generate password random
		password := utils.GenerateRandomPassword(10)
		hashed, _ := utils.HashPassword(password)
		// Buat akun user
		user := &rolesdomain.User{
			Email:    input.Patient.Email,
			Password: hashed,
			Role:     "pasien",
		}
		// Asumsikan ada userRepo di struct, jika tidak, tambahkan ke struct dan DI
		if ur, ok := u.userRepo.(userrepo.UserRepository); ok {
			ur.Create(ctx, user)
		}
		// Kirim email kredensial (log)
		subject := "Akun Klinik Anda"
		body := "Email: " + input.Patient.Email + "\nPassword: " + password
		utils.SendEmail(input.Patient.Email, subject, body)
		// Buat pasien
		u.patientRepo.Create(ctx, input.Patient)
	}
	// 2. Simpan screening (panggil repo screening answer)
	// 3. Masukkan ke antrian screening_pending (panggil repo queue)
	return nil
}
