package article

import "context"

type ArticleService interface {
	Load(ctx context.Context, id string) (*Article, error)
	Create(ctx context.Context, Article *Article) (int64, error)
	Update(ctx context.Context, Article *Article) (int64, error)
	Patch(ctx context.Context, Article map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}

func NewArticleService(repository ArticleRepository, generateId func(ctx context.Context) (string, error)) ArticleService {
	return &ArticleUseCase{repository: repository, generateId: generateId}
}

type ArticleUseCase struct {
	repository ArticleRepository
	generateId func(ctx context.Context) (string, error)
}

func (s *ArticleUseCase) Load(ctx context.Context, id string) (*Article, error) {
	return s.repository.Load(ctx, id)
}
func (s *ArticleUseCase) Create(ctx context.Context, article *Article) (int64, error) {
	id, err := s.generateId(ctx)
	if err != nil {
		return 0, err
	}
	article.Id = id
	return s.repository.Create(ctx, article)
}
func (s *ArticleUseCase) Update(ctx context.Context, article *Article) (int64, error) {
	return s.repository.Update(ctx, article)
}
func (s *ArticleUseCase) Patch(ctx context.Context, article map[string]interface{}) (int64, error) {
	return s.repository.Patch(ctx, article)
}
func (s *ArticleUseCase) Delete(ctx context.Context, id string) (int64, error) {
	return s.repository.Delete(ctx, id)
}
