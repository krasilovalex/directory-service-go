package usecase

import (
	"context"
	"directory-service/internal/domain"
	"time"

	"github.com/google/uuid"
)

type DepartmentUseCase struct {
	repo domain.DepartmentRepository
}

func NewDepartmentUseCase(r domain.DepartmentRepository) *DepartmentUseCase {
	return &DepartmentUseCase{
		repo: r,
	}
}

func (uc *DepartmentUseCase) Create(ctx context.Context, dept *domain.Department) error {
	dept.ID = uuid.New()
	dept.CreatedAt = time.Now()
	dept.IsActive = true

	return uc.repo.Create(ctx, dept)
}

func (uc *DepartmentUseCase) GetByID(ctx context.Context, id uuid.UUID) (*domain.Department, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *DepartmentUseCase) GetAll(ctx context.Context, limit, offset int) ([]domain.Department, error) {
	return uc.repo.GetAll(ctx, limit, offset)
}

func (uc *DepartmentUseCase) Update(ctx context.Context, dept *domain.Department) error {
	return uc.repo.Update(ctx, dept)
}

func (uc *DepartmentUseCase) Delete(ctx context.Context, id uuid.UUID) error {
	return uc.repo.Delete(ctx, id)
}
