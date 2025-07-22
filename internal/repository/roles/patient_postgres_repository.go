package roles

import (
	"context"
	"v2/internal/domain/roles"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PatientPostgresRepository struct {
	db *pgxpool.Pool
}

func NewPatientPostgresRepository(db *pgxpool.Pool) *PatientPostgresRepository {
	return &PatientPostgresRepository{db: db}
}

func (r *PatientPostgresRepository) Create(ctx context.Context, patient *roles.Patient) error {
	if patient.ID == uuid.Nil {
		patient.ID = uuid.New()
	}
	_, err := r.db.Exec(ctx, `INSERT INTO patients (
		id, nik, full_name, birth_place, birth_date, gender, address, rt, rw, village, district, religion, marital, job, nationality, valid_until, blood_type, height, weight, age, email, phone, ktp_images, created_at, updated_at
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25
	)`,
		patient.ID, patient.NIK, patient.FullName, patient.BirthPlace, patient.BirthDate, patient.Gender, patient.Address, patient.RT, patient.RW, patient.Village, patient.District, patient.Religion, patient.Marital, patient.Job, patient.Nationality, patient.ValidUntil, patient.BloodType, patient.Height, patient.Weight, patient.Age, patient.Email, patient.Phone, patient.KTPImages, patient.CreatedAt, patient.UpdatedAt)
	return err
}

func (r *PatientPostgresRepository) CreateOrUpdateByNIK(ctx context.Context, patient *roles.Patient) (*roles.Patient, error) {
	// Upsert by NIK
	query := `INSERT INTO patients (
		id, nik, full_name, birth_place, birth_date, gender, address, rt, rw, village, district, religion, marital, job, nationality, valid_until, blood_type, height, weight, age, email, phone, ktp_images, created_at, updated_at
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25
	)
	ON CONFLICT (nik) DO UPDATE SET
		full_name=EXCLUDED.full_name,
		birth_place=EXCLUDED.birth_place,
		birth_date=EXCLUDED.birth_date,
		gender=EXCLUDED.gender,
		address=EXCLUDED.address,
		rt=EXCLUDED.rt,
		rw=EXCLUDED.rw,
		village=EXCLUDED.village,
		district=EXCLUDED.district,
		religion=EXCLUDED.religion,
		marital=EXCLUDED.marital,
		job=EXCLUDED.job,
		nationality=EXCLUDED.nationality,
		valid_until=EXCLUDED.valid_until,
		blood_type=EXCLUDED.blood_type,
		height=EXCLUDED.height,
		weight=EXCLUDED.weight,
		age=EXCLUDED.age,
		email=EXCLUDED.email,
		phone=EXCLUDED.phone,
		ktp_images=EXCLUDED.ktp_images,
		updated_at=EXCLUDED.updated_at
	RETURNING id`
	if patient.ID == uuid.Nil {
		patient.ID = uuid.New()
	}
	row := r.db.QueryRow(ctx, query,
		patient.ID, patient.NIK, patient.FullName, patient.BirthPlace, patient.BirthDate, patient.Gender, patient.Address, patient.RT, patient.RW, patient.Village, patient.District, patient.Religion, patient.Marital, patient.Job, patient.Nationality, patient.ValidUntil, patient.BloodType, patient.Height, patient.Weight, patient.Age, patient.Email, patient.Phone, patient.KTPImages, patient.CreatedAt, patient.UpdatedAt)
	var id uuid.UUID
	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}
	patient.ID = id
	return patient, nil
}

func (r *PatientPostgresRepository) FindByNIK(ctx context.Context, nik string) (*roles.Patient, error) {
	row := r.db.QueryRow(ctx, `SELECT id, nik, full_name, birth_place, birth_date, gender, address, rt, rw, village, district, religion, marital, job, nationality, valid_until, blood_type, height, weight, age, email, phone, ktp_images, created_at, updated_at FROM patients WHERE nik=$1`, nik)
	return scanPatient(row)
}

func (r *PatientPostgresRepository) FindAll(ctx context.Context) ([]roles.Patient, error) {
	rows, err := r.db.Query(ctx, `SELECT id, nik, full_name, birth_place, birth_date, gender, address, rt, rw, village, district, religion, marital, job, nationality, valid_until, blood_type, height, weight, age, email, phone, ktp_images, created_at, updated_at FROM patients`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []roles.Patient
	for rows.Next() {
		p, err := scanPatient(rows)
		if err != nil {
			return nil, err
		}
		result = append(result, *p)
	}
	return result, nil
}

func (r *PatientPostgresRepository) FindAllPaginated(ctx context.Context, page, limit int) ([]roles.Patient, int64, error) {
	offset := (page - 1) * limit
	rows, err := r.db.Query(ctx, `SELECT id, nik, full_name, birth_place, birth_date, gender, address, rt, rw, village, district, religion, marital, job, nationality, valid_until, blood_type, height, weight, age, email, phone, ktp_images, created_at, updated_at FROM patients ORDER BY created_at DESC LIMIT $1 OFFSET $2`, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	var result []roles.Patient
	for rows.Next() {
		p, err := scanPatient(rows)
		if err != nil {
			return nil, 0, err
		}
		result = append(result, *p)
	}
	// Hitung total
	var total int64
	row := r.db.QueryRow(ctx, `SELECT COUNT(*) FROM patients`)
	if err := row.Scan(&total); err != nil {
		return nil, 0, err
	}
	return result, total, nil
}

func scanPatient(row interface {
	Scan(dest ...interface{}) error
}) (*roles.Patient, error) {
	var p roles.Patient
	var ktpImages []string
	err := row.Scan(
		&p.ID, &p.NIK, &p.FullName, &p.BirthPlace, &p.BirthDate, &p.Gender, &p.Address, &p.RT, &p.RW, &p.Village, &p.District, &p.Religion, &p.Marital, &p.Job, &p.Nationality, &p.ValidUntil, &p.BloodType, &p.Height, &p.Weight, &p.Age, &p.Email, &p.Phone, &ktpImages, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		return nil, err
	}
	p.KTPImages = ktpImages
	return &p, nil
}
