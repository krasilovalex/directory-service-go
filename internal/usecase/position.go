package usecase

import (
	"context"
	"directory-service/internal/domain"
	"time"

	"github.com/google/uuid"
)

type PositionUseCase struct {
	repo domain.PositionRepository
}

func NewPositionUseCase(r domain.PositionRepository) *PositionUseCase {
	return &PositionUseCase{
		repo: r,
	}
}

func (uc *PositionUseCase) Create(ctx context.Context, pos *domain.Position) error {
	pos.ID = uuid.New()
	pos.CreatedAt = time.Now()
	pos.IsActive = true

	return uc.repo.Create(ctx, pos)
}

func (uc *PositionUseCase) GetByID(ctx context.Context, id uuid.UUID) (*domain.Position, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *PositionUseCase) Update(ctx context.Context, pos *domain.Position) error {
	return uc.repo.Update(ctx, pos)
}

func (uc *PositionUseCase) Delete(ctx context.Context, id uuid.UUID) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *PositionUseCase) GetAll(ctx context.Context, limit, offset int) ([]domain.Position, error) {
	return uc.repo.GetAll(ctx, limit, offset)
}
