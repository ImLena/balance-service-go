package balance

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	AcceptPayment endpoint.Endpoint
	GetBalance    endpoint.Endpoint
	Receipt       endpoint.Endpoint
	Reserve       endpoint.Endpoint
	Report        endpoint.Endpoint
	Transactions  endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		AcceptPayment: makeAcceptPaymentEndpoint(s),
		GetBalance:    makeGetBalanceEndpoint(s),
		Receipt:       makeReceiptEndpoint(s),
		Reserve:       makeReserveEndpoint(s),
		Report:        makeReportEndpoint(s),
		Transactions:  makeTransactionsEndpoint(s),
	}
}

func makeGetBalanceEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetBalanceRequest)
		balance, err := s.GetBalance(req.Id)
		return GetBalanceResponse{Balance: balance}, err
	}
}

func makeReserveEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ReserveRequest)
		err := s.Reserve(req.UserID, req.ServiceID, req.OrderID, float32(req.Price), req.Comment)
		resp := ""
		if err == nil {
			resp = "Reservation successful"
		}
		return resp, err
	}
}

func makeReceiptEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ReceiptRequest)
		err := s.Receipt(req.UserID, req.Income, req.SourceID, req.Comment)
		resp := ""
		if err == nil {
			resp = "Successful"
		}
		return resp, err
	}
}

func makeAcceptPaymentEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ReserveRequest)
		err := s.AcceptPayment(req.UserID, req.ServiceID, req.OrderID, float32(req.Price))
		resp := ""
		if err == nil {
			resp = "Reservation verified"
		}
		return resp, err
	}
}

func makeReportEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ReportRequest)
		csv, err := s.Report(req.Year, req.Month)
		return csv, err
	}
}

func makeTransactionsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(TransactionsRequest)
		data, err := s.Transactions(req.UserID, req.Limit, req.Offset, req.Sort)
		return data, err
	}
}
