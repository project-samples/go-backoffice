package question

import (
	"context"
	sv "github.com/core-go/core"
)

type QuestionService interface {
	Load(ctx context.Context, id string) (*Question, error)
	Create(ctx context.Context, question *Question) (int64, error)
	Update(ctx context.Context, question *Question) (int64, error)
	Patch(ctx context.Context, question map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}

func NewQuestionService(repository sv.Repository) QuestionService {
	return &questionService{repository: repository}
}

type questionService struct {
	repository sv.Repository
}

func (s *questionService) Load(ctx context.Context, id string) (*Question, error) {
	var question Question
	ok, err := s.repository.LoadAndDecode(ctx, id, &question)
	if !ok {
		return nil, err
	} else {
		return &question, err
	}
}
func (s *questionService) Create(ctx context.Context, question *Question) (int64, error) {
	return s.repository.Insert(ctx, question)
}
func (s *questionService) Update(ctx context.Context, question *Question) (int64, error) {
	return s.repository.Update(ctx, question)
}
func (s *questionService) Patch(ctx context.Context, question map[string]interface{}) (int64, error) {
	return s.repository.Patch(ctx, question)
}
func (s *questionService) Delete(ctx context.Context, id string) (int64, error) {
	return s.repository.Delete(ctx, id)
}
