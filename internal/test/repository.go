package test

import "context"

type TestRepository interface {
	Load(ctx context.Context, id string) (*Test, error)
	Create(ctx context.Context, user *Test) (int64, error)
	Update(ctx context.Context, user *Test) (int64, error)
	Patch(ctx context.Context, user map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}
