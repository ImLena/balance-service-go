package balance

import "context"

type Receipt struct {
	ID       string  `json:"id,omitempty"`
	Income   float32 `json:"income"`
	SourceID int32   `json:"source_id"`
	UserID   string  `json:"user_id"`
	Comment  string  `json:"comment"`
}

type Reservation struct {
	ID        string  `json:"id,omitempty"`
	UserID    string  `json:"user_id"`
	ServiceID int32   `json:"service_id"`
	OrderID   int32   `json:"order_id"`
	Price     float32 `json:"price"`
	Verified  bool    `json:"verified"`
	Comment   string  `json:"comment"`
}

type Acceptation struct {
	UserID    string  `json:"user_id"`
	ServiceID int32   `json:"service_id"`
	OrderID   int32   `json:"order_id"`
	Price     float32 `json:"price"`
}

type Repository interface {
	AcceptPayment(ctx context.Context, acceptation Acceptation) error
	GetBalance(ctx context.Context, id string) (float32, error)
	Receipt(ctx context.Context, receipt Receipt) error
	Report(ctx context.Context, year string, month string) (map[int32]float64, error)
	Reserve(ctx context.Context, reservation Reservation) error
	Transactions(ctx context.Context, id string, limit int8, offset int8, sort string) ([]string, error)
}
