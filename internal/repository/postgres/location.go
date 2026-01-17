package postgres

import (
	"context"
	"directory-service/internal/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LocationRepository struct {
	db *pgxpool.Pool
}

func NewLocationRepository(db *pgxpool.Pool) *LocationRepository {
	return &LocationRepository{db: db}
}

func (r *LocationRepository) Create(ctx context.Context, loc *domain.Location) error {
	query := `
		INSERT INTO locations (id, name, is_active, created_at, address)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := r.db.Exec(ctx, query,
		loc.ID,
		loc.Name,
		loc.IsActive,
		loc.CreatedAt,
		loc.Address,
	)
	return err
}

func (r *LocationRepository) Update(ctx context.Context, loc *domain.Location) error {
	query := `UPDATE locations SET name = $2, is_active = $3, address = $4 WHERE id = $1`

	_, err := r.db.Exec(ctx, query,
		loc.ID,
		loc.Name,
		loc.IsActive,
		loc.Address,
	)

	return err
}

func (r *LocationRepository) GetByAddress(ctx context.Context, address string) (*domain.Location, error) {
	query := `SELECT id, name, is_active, created_at, address FROM locations WHERE address = $1`

	var loc domain.Location

	err := r.db.QueryRow(ctx, query, address).Scan(
		&loc.ID,
		&loc.Name,
		&loc.IsActive,
		&loc.CreatedAt,
		&loc.Address,
	)

	if err != nil {
		return nil, err
	}
	return &loc, nil
}

func (r *LocationRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Location, error) {
	query := `SELECT id, name, is_active, created_at, address FROM locations WHERE id = $1`

	var loc domain.Location

	err := r.db.QueryRow(ctx, query, id).Scan(
		&loc.ID,
		&loc.Name,
		&loc.IsActive,
		&loc.CreatedAt,
		&loc.Address,
	)

	if err != nil {
		return nil, err
	}
	return &loc, nil
}

func (r *LocationRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `UPDATE locations SET is_active = false WHERE id = $1`

	_, err := r.db.Exec(ctx, query, id)

	return err
}

func (r *LocationRepository) GetAll(ctx context.Context, limit, offset int) ([]domain.Location, error) {
	query := `
		SELECT id,name,is_active,created_at,address
		FROM locations
		ORDER BY name
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var locations []domain.Location

	for rows.Next() {
		var loc domain.Location

		if err := rows.Scan(&loc.ID, &loc.Name, &loc.IsActive, &loc.CreatedAt, &loc.Address); err != nil {
			return nil, err
		}

		locations = append(locations, loc)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return locations, nil
}
