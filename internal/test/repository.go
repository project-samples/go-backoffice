package test

import "context"

type TestRepository interface {
	Load(ctx context.Context, id string) (*Test, error)
	Create(ctx context.Context, test *Test) (int64, error)
	Update(ctx context.Context, test *Test) (int64, error)
	Patch(ctx context.Context, test map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}
