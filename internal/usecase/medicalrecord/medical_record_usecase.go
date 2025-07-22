package medicalrecord

import (
	"context"
	"fmt"
	"time"
	"v2/internal/domain/medicalrecord"
	repo "v2/internal/repository/medicalrecord"

	"github.com/google/uuid"
)

type MedicalRecordUsecase interface {
	CreateMedicalRecord(ctx context.Context, patientID string) (*medicalrecord.MedicalRecord, error)
}

type medicalRecordUsecase struct {
	recordRepo  repo.MedicalRecordRepository
	counterRepo repo.CounterRepository
}

func NewMedicalRecordUsecase(rr repo.MedicalRecordRepository, cr repo.CounterRepository) MedicalRecordUsecase {
	return &medicalRecordUsecase{
		recordRepo:  rr,
		counterRepo: cr,
	}
}

func (u *medicalRecordUsecase) CreateMedicalRecord(ctx context.Context, patientID string) (*medicalrecord.MedicalRecord, error) {
	// Cek apakah sudah ada MR untuk pasien ini
	pid, err := uuid.Parse(patientID)
	if err != nil {
		return nil, err
	}
	existing, err := u.recordRepo.FindByPatientID(ctx, pid)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return existing, nil // Sudah ada, return existing
	}
	// Ambil nomor urut berikutnya
	seq, err := u.counterRepo.GetNextSequence(ctx, "medical_record")
	if err != nil {
		return nil, err
	}
	mrNumber := fmt.Sprintf("MR%04d", seq)
	mr := &medicalrecord.MedicalRecord{
		PatientID: pid,
		MRNumber:  mrNumber,
		CreatedAt: time.Now(),
	}
	if err := u.recordRepo.Create(ctx, mr); err != nil {
		return nil, err
	}
	return mr, nil
}
