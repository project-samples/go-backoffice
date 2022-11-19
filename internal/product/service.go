package product

import "context"

type ProductService interface {
	Load(ctx context.Context, id string) (*Product, error)
	Create(ctx context.Context, Product *Product) (int64, error)
	Update(ctx context.Context, Product *Product) (int64, error)
	Patch(ctx context.Context, Product map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}

func NewProductService(repository ProductRepository, generateId func(ctx context.Context) (string, error)) ProductService {
	return &ProductUseCase{repository: repository, generateId: generateId}
}

type ProductUseCase struct {
	repository ProductRepository
	generateId func(ctx context.Context) (string, error)
}

func (s *ProductUseCase) Load(ctx context.Context, id string) (*Product, error) {
	return s.repository.Load(ctx, id)
}
func (s *ProductUseCase) Create(ctx context.Context, product *Product) (int64, error) {
	id, err := s.generateId(ctx)
	if err != nil {
		return 0, err
	}
	product.ProductId = id
	return s.repository.Create(ctx, product)
}
func (s *ProductUseCase) Update(ctx context.Context, product *Product) (int64, error) {
	return s.repository.Update(ctx, product)
}
func (s *ProductUseCase) Patch(ctx context.Context, product map[string]interface{}) (int64, error) {
	return s.repository.Patch(ctx, product)
}
func (s *ProductUseCase) Delete(ctx context.Context, id string) (int64, error) {
	return s.repository.Delete(ctx, id)
}
