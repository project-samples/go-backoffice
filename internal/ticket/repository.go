package ticket

import "context"

type TicketRepository interface {
	Load(ctx context.Context, id string) (*Ticket, error)
	Create(ctx context.Context, ticket *Ticket) (int64, error)
	Update(ctx context.Context, ticket *Ticket) (int64, error)
	Patch(ctx context.Context, ticket map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}
