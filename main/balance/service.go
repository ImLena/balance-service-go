package balance

import "context"

type Service interface {
	AcceptPayment(ctx context.Context, userID string, serviceID int64, orderID int64, price float32) error
	GetBalance(ctx context.Context, id string) (float32, error)
	Receipt(ctx context.Context, userID string, income float32) error
	Reserve(ctx context.Context, userID string, serviceID int64, orderID int64, price float32) error
}
