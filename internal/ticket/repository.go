package ticket

import "context"

type TicketRepository interface {
	Load(ctx context.Context, id string) (*Ticket, error)
	Create(ctx context.Context, user *Ticket) (int64, error)
	Update(ctx context.Context, user *Ticket) (int64, error)
	Patch(ctx context.Context, user map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}
