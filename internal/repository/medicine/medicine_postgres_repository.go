package medicine

import (
	"context"
	"v2/internal/domain/medicine"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MedicinePostgresRepository struct {
	db *pgxpool.Pool
}

func NewMedicinePostgresRepository(db *pgxpool.Pool) *MedicinePostgresRepository {
	return &MedicinePostgresRepository{db: db}
}

func (r *MedicinePostgresRepository) Create(ctx context.Context, m *medicine.Medicine) error {
	if m.ID == uuid.Nil {
		m.ID = uuid.New()
	}
	_, err := r.db.Exec(ctx, `INSERT INTO medicines (id, barcode, medicine_name, brand_name, category, dosage, content, quantity, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`, m.ID, m.Barcode, m.MedicineName, m.BrandName, m.Category, m.Dosage, m.Content, m.Quantity, m.CreatedAt, m.UpdatedAt)
	return err
}

func (r *MedicinePostgresRepository) Update(ctx context.Context, id uuid.UUID, update map[string]interface{}) error {
	_, err := r.db.Exec(ctx, `UPDATE medicines SET barcode=$1, medicine_name=$2, brand_name=$3, category=$4, dosage=$5, content=$6, quantity=$7, updated_at=NOW() WHERE id=$8`, update["barcode"], update["medicine_name"], update["brand_name"], update["category"], update["dosage"], update["content"], update["quantity"], id)
	return err
}

func (r *MedicinePostgresRepository) FindAll(ctx context.Context) ([]medicine.Medicine, error) {
	rows, err := r.db.Query(ctx, `SELECT id, barcode, medicine_name, brand_name, category, dosage, content, quantity, created_at, updated_at FROM medicines`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []medicine.Medicine
	for rows.Next() {
		var m medicine.Medicine
		if err := rows.Scan(&m.ID, &m.Barcode, &m.MedicineName, &m.BrandName, &m.Category, &m.Dosage, &m.Content, &m.Quantity, &m.CreatedAt, &m.UpdatedAt); err != nil {
			return nil, err
		}
		result = append(result, m)
	}
	return result, nil
}

func (r *MedicinePostgresRepository) FindAllPaginated(ctx context.Context, page, limit int) ([]medicine.Medicine, int64, error) {
	offset := (page - 1) * limit
	rows, err := r.db.Query(ctx, `SELECT id, barcode, medicine_name, brand_name, category, dosage, content, quantity, created_at, updated_at FROM medicines ORDER BY created_at DESC LIMIT $1 OFFSET $2`, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()
	var result []medicine.Medicine
	for rows.Next() {
		var m medicine.Medicine
		if err := rows.Scan(&m.ID, &m.Barcode, &m.MedicineName, &m.BrandName, &m.Category, &m.Dosage, &m.Content, &m.Quantity, &m.CreatedAt, &m.UpdatedAt); err != nil {
			return nil, 0, err
		}
		result = append(result, m)
	}
	// Hitung total
	var total int64
	row := r.db.QueryRow(ctx, `SELECT COUNT(*) FROM medicines`)
	if err := row.Scan(&total); err != nil {
		return nil, 0, err
	}
	return result, total, nil
}
