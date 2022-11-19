package question

import "context"

type QuestionRepository interface {
	Load(ctx context.Context, id string) (*Question, error)
	Create(ctx context.Context, user *Question) (int64, error)
	Update(ctx context.Context, user *Question) (int64, error)
	Patch(ctx context.Context, user map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}
