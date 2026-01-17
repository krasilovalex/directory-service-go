package postgres

import (
	"context"
	"directory-service/internal/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DepartmentRepository struct {
	db *pgxpool.Pool
}

func NewDepartmentRepository(db *pgxpool.Pool) *DepartmentRepository {
	return &DepartmentRepository{db: db}
}

func (r *DepartmentRepository) Create(ctx context.Context, dept *domain.Department) error {
	query := `
		INSERT INTO departments (id, name, is_active, created_at)
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.Exec(ctx, query,
		dept.ID,
		dept.Name,
		dept.IsActive,
		dept.CreatedAt,
	)

	return err
}

func (r *DepartmentRepository) Update(ctx context.Context, dept *domain.Department) error {
	query := `UPDATE departments SET name = $2, is_active = $3 WHERE id = $1 `

	_, err := r.db.Exec(ctx, query,
		dept.ID,
		dept.Name,
		dept.IsActive,
	)

	return err
}

func (r *DepartmentRepository) GetAll(ctx context.Context, limit, offset int) ([]domain.Department, error) {
	query := `
		SELECT id,name,is_active,created_at
		FROM departments
		ORDER BY name
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var departments []domain.Department

	for rows.Next() {
		var dept domain.Department

		if err := rows.Scan(&dept.ID, &dept.Name, &dept.IsActive, &dept.CreatedAt); err != nil {
			return nil, err
		}

		departments = append(departments, dept)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return departments, nil
}

func (r *DepartmentRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `UPDATE departments SET is_active = false WHERE id = $1`

	_, err := r.db.Exec(ctx, query, id)

	return err
}

func (r *DepartmentRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Department, error) {
	query := `SELECT id, name, is_active, created_at FROM departments WHERE id = $1`

	var dept domain.Department

	err := r.db.QueryRow(ctx, query, id).Scan(
		&dept.ID,
		&dept.Name,
		&dept.IsActive,
		&dept.CreatedAt,
	)

	if err != nil {
		return nil, err
	}
	return &dept, nil
}
