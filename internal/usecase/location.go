package usecase

import (
	"context"
	"directory-service/internal/domain"
	"time"

	"github.com/google/uuid"
)

type LocationUseCase struct {
	repo domain.LocationRepository
}

func NewLocationUseCase(r domain.LocationRepository) *LocationUseCase {
	return &LocationUseCase{
		repo: r,
	}
}

func (uc *LocationUseCase) Create(ctx context.Context, loc *domain.Location) error {
	loc.ID = uuid.New()
	loc.CreatedAt = time.Now()
	loc.IsActive = true

	return uc.repo.Create(ctx, loc)
}

func (uc *LocationUseCase) Update(ctx context.Context, loc *domain.Location) error {
	return uc.repo.Update(ctx, loc)
}

func (uc *LocationUseCase) GetByAddress(ctx context.Context, address string) (*domain.Location, error) {
	return uc.repo.GetByAddress(ctx, address)
}

func (uc *LocationUseCase) GetByID(ctx context.Context, id uuid.UUID) (*domain.Location, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *LocationUseCase) GetAll(ctx context.Context, limit, offset int) ([]domain.Location, error) {
	return uc.repo.GetAll(ctx, limit, offset)
}

func (uc *LocationUseCase) Delete(ctx context.Context, id uuid.UUID) error {
	return uc.repo.Delete(ctx, id)
}
