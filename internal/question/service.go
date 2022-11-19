package question

import "context"

type QuestionService interface {
	Load(ctx context.Context, id string) (*Question, error)
	GetQuestions(ctx context.Context, ids []string) ([]Question, error)
	Create(ctx context.Context, question *Question) (int64, error)
	Update(ctx context.Context, question *Question) (int64, error)
	Patch(ctx context.Context, question map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}

func NewQuestionService(repository QuestionRepository, generateId func(context.Context) (string, error)) QuestionService {
	return &QuestionUseCase{repository: repository, generateId: generateId}
}

type QuestionUseCase struct {
	repository QuestionRepository
	generateId func(context.Context) (string, error)
}

func (s *QuestionUseCase) GetQuestions(ctx context.Context, ids []string) ([]Question, error) {
	return s.repository.GetQuestions(ctx, ids)
}

func (s *QuestionUseCase) Load(ctx context.Context, id string) (*Question, error) {
	return s.repository.Load(ctx, id)
}
func (s *QuestionUseCase) Create(ctx context.Context, question *Question) (int64, error) {
	id, err := s.generateId(ctx)
	if err != nil {
		return 0, err
	}
	question.Id = id
	return s.repository.Create(ctx, question)
}
func (s *QuestionUseCase) Update(ctx context.Context, question *Question) (int64, error) {
	return s.repository.Update(ctx, question)
}
func (s *QuestionUseCase) Patch(ctx context.Context, question map[string]interface{}) (int64, error) {
	return s.repository.Patch(ctx, question)
}
func (s *QuestionUseCase) Delete(ctx context.Context, id string) (int64, error) {
	return s.repository.Delete(ctx, id)
}
