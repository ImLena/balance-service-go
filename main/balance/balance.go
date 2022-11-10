package balance

import "context"

type Receipt struct {
	ID     string  `json:"id,omitempty"`
	Income float32 `json:"income"`
}

type Reservation struct {
	ID        string  `json:"id,omitempty"`
	UserID    string  `json:"user_id"`
	ServiceID int64   `json:"service_id"`
	OrderID   int64   `json:"order_id"`
	Price     float32 `json:"price"`
	Verified  bool    `json:"verified"`
}

type Acceptation struct {
	UserID    string  `json:"user_id"`
	ServiceID int64   `json:"service_id"`
	OrderID   int64   `json:"order_id"`
	Price     float32 `json:"price"`
}

type Repository interface {
	AcceptPayment(ctx context.Context, acceptation Acceptation) error
	GetBalance(ctx context.Context, id string) (float32, error)
	Receipt(ctx context.Context, receipt Receipt) error
	Reserve(ctx context.Context, reservation Reservation) error
}
