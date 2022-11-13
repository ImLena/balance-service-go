package balance

type Service interface {
	AcceptPayment(userID string, serviceID int32, orderID int32, price float32) error
	GetBalance(id string) (float32, error)
	Receipt(userID string, income float32, sourceID int32, comment string) error
	Transactions(userID string, limit int8, offset int8, sort string) ([]string, error)
	Report(year string, month string) (string, error)
	Reserve(userID string, serviceID int32, orderID int32, price float32, comment string) error
}
