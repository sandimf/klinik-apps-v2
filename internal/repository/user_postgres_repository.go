package repository

import (
	"context"
	roles "v2/internal/domain/roles"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserPostgresRepository struct {
	db *pgxpool.Pool
}

func NewUserPostgresRepository(db *pgxpool.Pool) *UserPostgresRepository {
	return &UserPostgresRepository{db: db}
}

func (r *UserPostgresRepository) Create(ctx context.Context, user *roles.User) (string, error) {
	if user.ID == uuid.Nil {
		user.ID = uuid.New()
	}
	_, err := r.db.Exec(ctx, `INSERT INTO users (id, email, password, role) VALUES ($1, $2, $3, $4)`,
		user.ID, user.Email, user.Password, user.Role)
	if err != nil {
		return "", err
	}
	return user.ID.String(), nil
}

func (r *UserPostgresRepository) FindByEmail(ctx context.Context, email string) (*roles.User, error) {
	row := r.db.QueryRow(ctx, `SELECT id, email, password, role FROM users WHERE email=$1`, email)
	var user roles.User
	var id uuid.UUID
	if err := row.Scan(&id, &user.Email, &user.Password, &user.Role); err != nil {
		return nil, err
	}
	user.ID = id
	return &user, nil
}

func (r *UserPostgresRepository) FindByID(ctx context.Context, id string) (*roles.User, error) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	row := r.db.QueryRow(ctx, `SELECT id, email, password, role FROM users WHERE id=$1`, uuidID)
	var user roles.User
	var uid uuid.UUID
	if err := row.Scan(&uid, &user.Email, &user.Password, &user.Role); err != nil {
		return nil, err
	}
	user.ID = uid
	return &user, nil
}
