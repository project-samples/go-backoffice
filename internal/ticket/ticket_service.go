package ticket

import (
	"context"
	sv "github.com/core-go/core"
)

type TicketService interface {
	Load(ctx context.Context, id string) (*Ticket, error)
	Create(ctx context.Context, ticket *Ticket) (int64, error)
	Update(ctx context.Context, ticket *Ticket) (int64, error)
	Patch(ctx context.Context, ticket map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}

func NewTicketService(repository sv.Repository) TicketService {
	return &TicketUseCase{repository: repository}
}

type TicketUseCase struct {
	repository sv.Repository
}

func (s *TicketUseCase) Load(ctx context.Context, id string) (*Ticket, error) {
	var ticket Ticket
	ok, err := s.repository.LoadAndDecode(ctx, id, &ticket)
	if !ok {
		return nil, err
	} else {
		return &ticket, err
	}
}
func (s *TicketUseCase) Create(ctx context.Context, ticket *Ticket) (int64, error) {
	return s.repository.Insert(ctx, ticket)
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
