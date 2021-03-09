package persistance

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type Order struct {
	Id                 uuid.UUID
	MenuItems          []MenuItem
	OrderedAtTimestamp time.Time
	Cost               int
}

type MenuItem struct {
	Id       uuid.UUID
	Quantity int
}

type OrderRepository interface {
	Get(ctx context.Context, uuid uuid.UUID) (Order, error)
	GetList(ctx context.Context, offset int, amount int) ([]Order, error)
	Create(ctx context.Context, order Order) error
}
