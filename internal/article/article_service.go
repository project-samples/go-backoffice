package article

import (
	"context"
	sv "github.com/core-go/core"
)

type ArticleService interface {
	Load(ctx context.Context, id string) (*Article, error)
	Create(ctx context.Context, Article *Article) (int64, error)
	Update(ctx context.Context, Article *Article) (int64, error)
	Patch(ctx context.Context, Article map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}

func NewArticleService(repository sv.Repository) ArticleService {
	return &ArticleUseCase{repository: repository}
}

type ArticleUseCase struct {
	repository sv.Repository
}

func (s *ArticleUseCase) Load(ctx context.Context, id string) (*Article, error) {
	var Article Article
	ok, err := s.repository.LoadAndDecode(ctx, id, &Article)
	if !ok {
		return nil, err
	} else {
		return &Article, err
	}
}
func (s *ArticleUseCase) Create(ctx context.Context, Article *Article) (int64, error) {
	return s.repository.Insert(ctx, Article)
}
func (s *ArticleUseCase) Update(ctx context.Context, Article *Article) (int64, error) {
	return s.repository.Update(ctx, Article)
}
func (s *ArticleUseCase) Patch(ctx context.Context, Article map[string]interface{}) (int64, error) {
	return s.repository.Patch(ctx, Article)
}
func (s *ArticleUseCase) Delete(ctx context.Context, id string) (int64, error) {
	return s.repository.Delete(ctx, id)
}
