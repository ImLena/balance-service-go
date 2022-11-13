package balance

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
	AcceptPayment(acceptation Acceptation) error
	GetBalance(id string) (float32, error)
	Receipt(receipt Receipt) error
	Report(year string, month string) (map[int32]float64, error)
	Reserve(reservation Reservation) error
	Transactions(id string, limit int8, offset int8, sort string) ([]string, error)
}
