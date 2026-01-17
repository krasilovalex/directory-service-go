package postgres

import (
	"context"
	"directory-service/internal/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PositionRepository struct {
	db *pgxpool.Pool
}

func NewPositionRepository(db *pgxpool.Pool) *PositionRepository {
	return &PositionRepository{db: db}
}

func (r *PositionRepository) Create(ctx context.Context, pos *domain.Position) error {
	query := `
		INSERT INTO positions (id, name, is_active, created_at)
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.Exec(ctx, query,
		pos.ID,
		pos.Name,
		pos.IsActive,
		pos.CreatedAt,
	)
	return err
}

func (r *PositionRepository) Update(ctx context.Context, pos *domain.Position) error {
	query := `UPDATE positions SET name = $2, is_active = $3  WHERE id = $1`

	_, err := r.db.Exec(ctx, query,
		pos.ID,
		pos.Name,
		pos.IsActive,
	)

	return err
}

func (r *PositionRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Position, error) {
	query := `SELECT id, name, is_active, created_at FROM positions WHERE id = $1`

	var pos domain.Position

	err := r.db.QueryRow(ctx, query, id).Scan(
		&pos.ID,
		&pos.Name,
		&pos.IsActive,
		&pos.CreatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &pos, nil
}

func (r *PositionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `UPDATE positions SET is_active = false WHERE id = $1`

	_, err := r.db.Exec(ctx, query, id)

	return err
}

func (r *PositionRepository) GetAll(ctx context.Context, limit, offset int) ([]domain.Position, error) {
	query := `
		SELECT id,name,is_active,created_at
		FROM positions
		ORDER BY name
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var positions []domain.Position

	for rows.Next() {
		var pos domain.Position

		if err := rows.Scan(&pos.ID, &pos.Name, &pos.IsActive, &pos.CreatedAt); err != nil {
			return nil, err
		}

		positions = append(positions, pos)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return positions, nil
}
