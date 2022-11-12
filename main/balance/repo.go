package balance

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

var NegativeBalance = errors.New("Insufficient funds in the account")
var AcceptationNotFoundErr = errors.New("Unable to find reservation")
var AcceptationErr = errors.New("Reservation already verified")

type repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) Repository {
	return &repo{
		db: db,
	}
}

func (repo *repo) Reserve(ctx context.Context, reservation Reservation) error {
	tx, err := repo.db.Begin()
	var balance float32
	err = tx.QueryRow("SELECT balance FROM balance WHERE id=$1",
		reservation.UserID).Scan(&balance)

	newBalance := balance - reservation.Price
	if newBalance < 0 {
		tx.Rollback()
		return NegativeBalance
	}
	_, err = tx.Exec("UPDATE balance SET balance=$2 WHERE id=$1",
		reservation.UserID, newBalance)

	_, err = tx.Exec("INSERT INTO transactions (id, service_id, order_id, price, user_id, verified, comment, time) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		reservation.ID, reservation.ServiceID, reservation.OrderID, -reservation.Price,
		reservation.UserID, reservation.Verified, reservation.Comment, time.Now())
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()

	return err
}

func (repo *repo) AcceptPayment(ctx context.Context, acceptation Acceptation) error {
	tx, err := repo.db.Begin()
	var id string
	var verify bool
	err = tx.QueryRow("SELECT id, verified FROM transactions WHERE user_id=$1 AND service_id=$2 AND order_id=$3 AND price=$4",
		acceptation.UserID, acceptation.ServiceID, acceptation.OrderID, -acceptation.Price).Scan(&id, &verify)
	if err != nil {
		tx.Rollback()
		return AcceptationNotFoundErr
	}
	if verify {
		tx.Rollback()
		return AcceptationErr
	}
	_, err = tx.Exec("UPDATE transactions SET verified=true WHERE id=$1", id)

	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()

	return err
}

func (repo *repo) Receipt(ctx context.Context, receipt Receipt) error {
	tx, err := repo.db.Begin()
	_, err = tx.Exec("INSERT INTO balance (id, balance) VALUES ($1, $2)",
		receipt.UserID, receipt.Income)

	_, err = tx.Exec("INSERT INTO transactions (id, service_id, order_id, price, user_id, verified, comment, time) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		receipt.ID, receipt.SourceID, nil, receipt.Income,
		receipt.UserID, true, receipt.Comment, time.Now())

	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()

	return nil
}

func (repo *repo) GetBalance(ctx context.Context, id string) (float32, error) {
	var balance float32
	err := repo.db.QueryRow("SELECT balance FROM balance WHERE id=$1", id).Scan(&balance)
	if err != nil {
		return -1, err
	}

	return balance, nil
}

func (repo *repo) Transactions(ctx context.Context, id string, limit int8, offset int8, sort string) ([]string, error) {
	rows, err := repo.db.Query("SELECT service_id, price, comment, time FROM transactions WHERE user_id=$1 ORDER BY $4 LIMIT $2 OFFSET $3",
		id, limit, offset, sort)
	if err != nil {
		return nil, err
	}

	var data []string
	for rows.Next() {
		var serviceId int32
		var price int32
		var comment string
		var t time.Time
		err = rows.Scan(&serviceId, &price, &comment, &t)

		if err != nil {
			return nil, err
		}

		if price > 0 {
			data = append(data, fmt.Sprintf("Income %d from %d with comment: '%v' ", price, serviceId, comment))
		} else {
			data = append(data, fmt.Sprintf("Debited %d to %d with comment: '%v' ", -price, serviceId, comment))
		}
	}

	return data, nil
}

func (repo *repo) Report(ctx context.Context, year string, month string) (map[int32]float64, error) {
	rows, err := repo.db.Query("SELECT service_id, price FROM transactions WHERE EXTRACT(YEAR FROM time)=$1 AND EXTRACT(MONTH FROM time)=$2 AND verified=true AND price<0",
		year, month)

	if err != nil {
		return nil, err
	}

	data := make(map[int32]float64)
	for rows.Next() {
		var serviceId int32
		var price float64
		err = rows.Scan(&serviceId, &price)

		if err != nil {
			return nil, err
		}
		data[serviceId] += price
	}

	return data, nil
}
