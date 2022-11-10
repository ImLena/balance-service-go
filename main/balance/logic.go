package balance

import (
	"context"
	"github.com/gofrs/uuid"

	"github.com/go-kit/kit/log"
)

type service struct {
	repository Repository
	logger     log.Logger
}

func (s service) GetBalance(ctx context.Context, id string) (float32, error) {
	balance, err := s.repository.GetBalance(ctx, id)

	if err != nil {
		//level.Error(logger).Log("err", err)
		return -1, err
	}

	return balance, err
}

func (s service) Reserve(ctx context.Context, userID string, serviceID int64, orderID int64, price float32) error {

	id, _ := uuid.NewV4()
	reservationID := id.String()
	err := s.repository.Reserve(ctx, Reservation{reservationID, userID, serviceID, orderID, price, false})

	return err
}

func (s service) AcceptPayment(ctx context.Context, userID string, serviceID int64, orderID int64, price float32) error {
	err := s.repository.AcceptPayment(ctx, Acceptation{userID, serviceID, orderID, price})

	return err
}

func (s service) Receipt(ctx context.Context, userID string, income float32) error {
	err := s.repository.Receipt(ctx, Receipt{userID, income})

	return err
}

func NewService(rep Repository, logger log.Logger) Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}
