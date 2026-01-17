package domain

import (
	"context"

	"github.com/google/uuid"
)

type DepartmentRepository interface {
	Create(ctx context.Context, dept *Department) error
	GetByID(ctx context.Context, id uuid.UUID) (*Department, error)
	GetAll(ctx context.Context, limit, offset int) ([]Department, error)
	Update(ctx context.Context, dept *Department) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type PositionRepository interface {
	Create(ctx context.Context, dept *Position) error
	GetByID(ctx context.Context, id uuid.UUID) (*Position, error)
	GetAll(ctx context.Context, limit, offset int) ([]Position, error)
	Update(ctx context.Context, dept *Position) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type LocationRepository interface {
	Create(ctx context.Context, dept *Location) error
	GetByAddress(ctx context.Context, address string) (*Location, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Location, error)
	GetAll(ctx context.Context, limit, offset int) ([]Location, error)
	Update(ctx context.Context, dept *Location) error
	Delete(ctx context.Context, id uuid.UUID) error
}
