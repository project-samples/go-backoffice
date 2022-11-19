package term

import "context"

type TermService interface {
	Load(ctx context.Context, id string) (*Term, error)
	Create(ctx context.Context, Term *Term) (int64, error)
	Update(ctx context.Context, Term *Term) (int64, error)
	Patch(ctx context.Context, Term map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}

func NewTermService(repository TermRepository, generateId func(ctx context.Context) (string, error)) TermService {
	return &TermUseCase{repository: repository, generateId: generateId}
}

type TermUseCase struct {
	repository TermRepository
	generateId func(ctx context.Context) (string, error)
}

func (s *TermUseCase) Load(ctx context.Context, id string) (*Term, error) {
	return s.repository.Load(ctx, id)
}
func (s *TermUseCase) Create(ctx context.Context, term *Term) (int64, error) {
	id, err := s.generateId(ctx)
	if err != nil {
		return 0, err
	}
	term.Id = id
	return s.repository.Create(ctx, term)
}
func (s *TermUseCase) Update(ctx context.Context, term *Term) (int64, error) {
	return s.repository.Update(ctx, term)
}
func (s *TermUseCase) Patch(ctx context.Context, term map[string]interface{}) (int64, error) {
	return s.repository.Patch(ctx, term)
}
func (s *TermUseCase) Delete(ctx context.Context, id string) (int64, error) {
	return s.repository.Delete(ctx, id)
}
