package company

import "context"

type CompanyRepository interface {
	Load(ctx context.Context, id string) (*Company, error)
	Create(ctx context.Context, company *Company) (int64, error)
	Update(ctx context.Context, company *Company) (int64, error)
	Patch(ctx context.Context, company map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}
