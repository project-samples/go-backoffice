package ticket

import "context"

type TicketService interface {
	Load(ctx context.Context, id string) (*Ticket, error)
	Create(ctx context.Context, ticket *Ticket) (int64, error)
	Update(ctx context.Context, ticket *Ticket) (int64, error)
	Patch(ctx context.Context, ticket map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}

func NewTicketService(repository TicketRepository, generateId func(ctx context.Context) (string, error)) TicketService {
	return &TicketUseCase{repository: repository, generateId: generateId}
}

type TicketUseCase struct {
	repository TicketRepository
	generateId func(ctx context.Context) (string, error)
}

func (s *TicketUseCase) Load(ctx context.Context, id string) (*Ticket, error) {
	return s.repository.Load(ctx, id)
}
func (s *TicketUseCase) Create(ctx context.Context, ticket *Ticket) (int64, error) {
	id, err := s.generateId(ctx)
	if err != nil {
		return 0, err
	}
	ticket.Id = id
	return s.repository.Create(ctx, ticket)
}
func (s *TicketUseCase) Update(ctx context.Context, ticket *Ticket) (int64, error) {
	return s.repository.Update(ctx, ticket)
}
func (s *TicketUseCase) Patch(ctx context.Context, ticket map[string]interface{}) (int64, error) {
	return s.repository.Patch(ctx, ticket)
}
func (s *TicketUseCase) Delete(ctx context.Context, id string) (int64, error) {
	return s.repository.Delete(ctx, id)
}
