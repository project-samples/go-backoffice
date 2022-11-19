package term

import "context"

type TermRepository interface {
	Load(ctx context.Context, id string) (*Term, error)
	Create(ctx context.Context, term *Term) (int64, error)
	Update(ctx context.Context, term *Term) (int64, error)
	Patch(ctx context.Context, term map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}
