package test

import (
	"context"
	sv "github.com/core-go/core"
)

type TestService interface {
	Load(ctx context.Context, id string) (*Test, error)
	Create(ctx context.Context, test *Test) (int64, error)
	Update(ctx context.Context, test *Test) (int64, error)
	Patch(ctx context.Context, test map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}

func NewTestService(repository sv.Repository) TestService {
	return &TestUseCase{repository: repository}
}

type TestUseCase struct {
	repository sv.Repository
}

func (s *TestUseCase) Load(ctx context.Context, id string) (*Test, error) {
	var test Test
	ok, err := s.repository.LoadAndDecode(ctx, id, &test)
	if !ok {
		return nil, err
	} else {
		return &test, err
	}
}
func (s *TestUseCase) Create(ctx context.Context, test *Test) (int64, error) {
	return s.repository.Insert(ctx, test)
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
