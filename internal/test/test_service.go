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
	return &testService{repository: repository}
}

type testService struct {
	repository sv.Repository
}

func (s *testService) Load(ctx context.Context, id string) (*Test, error) {
	var test Test
	ok, err := s.repository.LoadAndDecode(ctx, id, &test)
	if !ok {
		return nil, err
	} else {
		return &test, err
	}
}
func (s *testService) Create(ctx context.Context, test *Test) (int64, error) {
	return s.repository.Insert(ctx, test)
}
func (s *testService) Update(ctx context.Context, test *Test) (int64, error) {
	return s.repository.Update(ctx, test)
}
func (s *testService) Patch(ctx context.Context, test map[string]interface{}) (int64, error) {
	return s.repository.Patch(ctx, test)
}
func (s *testService) Delete(ctx context.Context, id string) (int64, error) {
	return s.repository.Delete(ctx, id)
}
