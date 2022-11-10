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
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		AcceptPayment: makeAcceptPaymentEndpoint(s),
		GetBalance:    makeGetBalanceEndpoint(s),
		Receipt:       makeReceiptEndpoint(s),
		Reserve:       makeReserveEndpoint(s),
	}
}

func makeGetBalanceEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetBalanceRequest)
		balance, err := s.GetBalance(ctx, req.Id)
		return GetBalanceResponse{Balance: balance}, err
	}
}

func makeReserveEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ReserveRequest)
		err := s.Reserve(ctx, req.UserID, req.ServiceID, req.OrderID, float32(req.Price))
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
		err := s.Receipt(ctx, req.UserID, req.Income)
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
		err := s.AcceptPayment(ctx, req.UserID, req.ServiceID, req.OrderID, float32(req.Price))
		resp := ""
		if err == nil {
			resp = "Reservation verified"
		}
		return resp, err
	}
}
