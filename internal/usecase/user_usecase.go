package usecase

import (
	"context"
	"errors"
	"time"
	"v2/internal/domain/roles"
	userrepo "v2/internal/repository"
	patientrepo "v2/internal/repository/roles"
	"v2/internal/utils"
)

type RegisterPatientInput struct {
	NIK         string
	FullName    string
	BirthPlace  string
	BirthDate   string
	Gender      string
	Address     string
	RT          string
	RW          string
	Village     string
	District    string
	Religion    string
	Marital     string
	Job         string
	Nationality string
	ValidUntil  string
	BloodType   string
	Height      int
	Weight      int
	Age         int
	Email       string
	Phone       string
	Password    string
	KTPImages   []string
}

type UserUsecase interface {
	RegisterPatient(ctx context.Context, input RegisterPatientInput) error
	Login(ctx context.Context, email, password string) (string, error)
}

type userUsecase struct {
	userRepo    userrepo.UserRepository
	patientRepo patientrepo.PatientRepository
}

func NewUserUsecase(ur userrepo.UserRepository, pr patientrepo.PatientRepository) UserUsecase {
	return &userUsecase{
		userRepo:    ur,
		patientRepo: pr,
	}
}

func (uc *userUsecase) RegisterPatient(ctx context.Context, input RegisterPatientInput) error {
	// 1. Check if email exists
	existingUser, err := uc.userRepo.FindByEmail(ctx, input.Email)
	if err != nil {
		return err // Internal server error
	}
	if existingUser != nil {
		return errors.New("email already exists")
	}

	// 2. Hash password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return err
	}

	// 3. Create User
	user := &roles.User{
		Email:    input.Email,
		Password: hashedPassword,
		Role:     "pasien",
	}

	_, err = uc.userRepo.Create(ctx, user)
	if err != nil {
		return err
	}

	// 4. Create or update Patient Profile
	patient := &roles.Patient{
		NIK:         input.NIK,
		FullName:    input.FullName,
		BirthPlace:  input.BirthPlace,
		BirthDate:   input.BirthDate,
		Gender:      input.Gender,
		Address:     input.Address,
		RT:          input.RT,
		RW:          input.RW,
		Village:     input.Village,
		District:    input.District,
		Religion:    input.Religion,
		Marital:     input.Marital,
		Job:         input.Job,
		Nationality: input.Nationality,
		ValidUntil:  input.ValidUntil,
		BloodType:   input.BloodType,
		Height:      input.Height,
		Weight:      input.Weight,
		Age:         input.Age,
		Email:       input.Email,
		Phone:       input.Phone,
		KTPImages:   input.KTPImages,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	_, err = uc.patientRepo.CreateOrUpdateByNIK(ctx, patient)
	return err
}

func (uc *userUsecase) Login(ctx context.Context, email, password string) (string, error) {
	user, err := uc.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("invalid email or password")
	}
	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid email or password")
	}

	token, err := utils.GenerateJWT(user.ID.String(), user.Role, time.Hour*1) // 1 jam
	if err != nil {
		return "", err
	}
	return token, nil
}
