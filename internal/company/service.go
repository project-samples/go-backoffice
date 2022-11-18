package company

import "context"

type CompanyService interface {
	Load(ctx context.Context, id string) (*Company, error)
	Create(ctx context.Context, company *Company) (int64, error)
	Update(ctx context.Context, company *Company) (int64, error)
	Patch(ctx context.Context, company map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}

func NewCompanyService(repository CompanyRepository) CompanyService {
	return &CompanyUseCase{repository: repository}
}

type CompanyUseCase struct {
	repository CompanyRepository
}

func (s *CompanyUseCase) Load(ctx context.Context, id string) (*Company, error) {
	return s.repository.Load(ctx, id)
}
func (s *CompanyUseCase) Create(ctx context.Context, company *Company) (int64, error) {
	return s.repository.Create(ctx, company)
}
func (s *CompanyUseCase) Update(ctx context.Context, company *Company) (int64, error) {
	return s.repository.Update(ctx, company)
}
func (s *CompanyUseCase) Patch(ctx context.Context, company map[string]interface{}) (int64, error) {
	return s.repository.Patch(ctx, company)
}
func (s *CompanyUseCase) Delete(ctx context.Context, id string) (int64, error) {
	return s.repository.Delete(ctx, id)
}
