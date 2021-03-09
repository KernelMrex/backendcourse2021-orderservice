package persistance

import (
	"context"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type MysqlOrderRepository struct {
	db *sqlx.DB
}

func NewMysqlOrderRepository(dataSourceName string) (OrderRepository, error) {
	db, err := sqlx.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &MysqlOrderRepository{db: db}, nil
}

func (r *MysqlOrderRepository) Get(ctx context.Context, uuid uuid.UUID) (Order, error) {
	return Order{}, nil
}

func (r *MysqlOrderRepository) GetList(ctx context.Context, offset int, amount int) ([]Order, error) {
	return []Order{}, nil
}

func (r *MysqlOrderRepository) Create(ctx context.Context, order Order) error {
	return nil
}
