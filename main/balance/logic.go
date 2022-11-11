package balance

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/gofrs/uuid"
	"log"
	"os"
)

type service struct {
	repository Repository
}

func (s service) GetBalance(ctx context.Context, id string) (float32, error) {
	balance, err := s.repository.GetBalance(ctx, id)

	if err != nil {
		//level.Error(logger).Log("err", err)
		return -1, err
	}

	return balance, err
}

func (s service) Reserve(ctx context.Context, userID string, serviceID int32, orderID int32, price float32, comment string) error {

	id, _ := uuid.NewV4()
	reservationID := id.String()
	err := s.repository.Reserve(ctx, Reservation{reservationID, userID, serviceID, orderID, price, false, comment})

	return err
}

func (s service) AcceptPayment(ctx context.Context, userID string, serviceID int32, orderID int32, price float32) error {
	err := s.repository.AcceptPayment(ctx, Acceptation{userID, serviceID, orderID, price})

	return err
}

func (s service) Receipt(ctx context.Context, userID string, income float32, comment string) error {
	err := s.repository.Receipt(ctx, Receipt{userID, income, comment})

	return err
}

func (s service) Report(ctx context.Context, year string, month string) (string, error) {
	report, err := s.repository.Report(ctx, year, month)
	name := "reports/" + year + "-" + month + ".csv"
	f, err := os.Create(name)

	if err != nil {
		log.Fatal("Can not create file", err)
	}

	w := csv.NewWriter(f)
	w.Comma = ';'
	for key, value := range report {
		str := []string{fmt.Sprintf("%v", key), fmt.Sprintf("%v", value)}
		err := w.Write(str)
		if err != nil {
			log.Fatal(err)
		}
	}
	w.Flush()

	return name, err
}

func NewService(rep Repository) Service {
	return &service{
		repository: rep,
	}
}
