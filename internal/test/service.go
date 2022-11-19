package test

import "context"

type TestService interface {
	Load(ctx context.Context, id string) (*Test, error)
	Create(ctx context.Context, test *Test) (int64, error)
	Update(ctx context.Context, test *Test) (int64, error)
	Patch(ctx context.Context, test map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}

func NewTestService(repository TestRepository) TestService {
	return &TestUseCase{repository: repository}
}

type TestUseCase struct {
	repository TestRepository
}

func (s *TestUseCase) Load(ctx context.Context, id string) (*Test, error) {
	return s.repository.Load(ctx, id)
}
func (s *TestUseCase) Create(ctx context.Context, test *Test) (int64, error) {
	return s.repository.Create(ctx, test)
}
func (s *TestUseCase) Update(ctx context.Context, test *Test) (int64, error) {
	return s.repository.Update(ctx, test)
}
func (s *TestUseCase) Patch(ctx context.Context, test map[string]interface{}) (int64, error) {
	return s.repository.Patch(ctx, test)
}
func (s *TestUseCase) Delete(ctx context.Context, id string) (int64, error) {
	return s.repository.Delete(ctx, id)
}
