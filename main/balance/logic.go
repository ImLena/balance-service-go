package balance

import (
	"encoding/csv"
	"fmt"
	"github.com/gofrs/uuid"
	"log"
	"os"
)

type service struct {
	repository Repository
}

func (s service) GetBalance(id string) (float32, error) {
	balance, err := s.repository.GetBalance(id)

	if err != nil {
		//level.Error(logger).Log("err", err)
		return -1, err
	}

	return balance, err
}

func (s service) Reserve(userID string, serviceID int32, orderID int32, price float32, comment string) error {

	id, _ := uuid.NewV4()
	reservationID := id.String()
	err := s.repository.Reserve(Reservation{reservationID, userID, serviceID, orderID, price, false, comment})

	return err
}

func (s service) AcceptPayment(userID string, serviceID int32, orderID int32, price float32) error {
	err := s.repository.AcceptPayment(Acceptation{userID, serviceID, orderID, price})

	return err
}

func (s service) Receipt(userID string, income float32, sourceID int32, comment string) error {
	id, _ := uuid.NewV4()
	receiptID := id.String()
	err := s.repository.Receipt(Receipt{receiptID, income, sourceID, userID, comment})

	return err
}

func (s service) Transactions(userID string, limit int8, offset int8, sort string) ([]string, error) {
	data, err := s.repository.Transactions(userID, limit, offset, sort)
	return data, err
}

func (s service) Report(year string, month string) (string, error) {
	report, err := s.repository.Report(year, month)
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
