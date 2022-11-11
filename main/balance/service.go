package balance

import "context"

type Service interface {
	AcceptPayment(ctx context.Context, userID string, serviceID int32, orderID int32, price float32) error
	GetBalance(ctx context.Context, id string) (float32, error)
	Receipt(ctx context.Context, userID string, income float32, comment string) error
	Report(ctx context.Context, year string, month string) (string, error)
	Reserve(ctx context.Context, userID string, serviceID int32, orderID int32, price float32, comment string) error
}
