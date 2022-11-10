package balance

import (
	"context"
	"database/sql"
	"errors"
)

var RepoErr = errors.New("Unable to handle Repo Request")
var NegativeBalance = errors.New("Insufficient funds in the account")
var AcceptationErr = errors.New("Unable to find reservation")

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

	_, err = tx.Exec("INSERT INTO transactions (id, service_id, order_id, price, user_id, verified) VALUES ($1, $2, $3, $4, $5, $6)",
		reservation.ID, reservation.ServiceID, reservation.OrderID, reservation.Price, reservation.UserID, reservation.Verified)
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
	err = tx.QueryRow("SELECT id FROM transactions WHERE user_id=$1 AND service_id=$2 AND order_id=$3 AND price=$4",
		acceptation.UserID, acceptation.ServiceID, acceptation.OrderID, acceptation.Price).Scan(&id)
	if err != nil {
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
	_, err := repo.db.Exec("INSERT INTO balance (id, balance) VALUES ($1, $2)",
		receipt.ID, receipt.Income)
	if err != nil {
		return err
	}
	return nil
}

func (repo *repo) GetBalance(ctx context.Context, id string) (float32, error) {
	var balance float32
	err := repo.db.QueryRow("SELECT balance FROM balance WHERE id=$1", id).Scan(&balance)
	if err != nil {
		return -1, RepoErr

	}

	return balance, nil
}
