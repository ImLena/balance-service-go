package balance

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type (
	GetBalanceRequest struct {
		Id string `json:"id"`
	}
	GetBalanceResponse struct {
		Balance float32 `json:"balance"`
	}
	ReceiptRequest struct {
		UserID  string  `json:"user_id"`
		Income  float32 `json:"income"`
		Comment string  `json:"comment"`
	}
	ReserveRequest struct {
		UserID    string  `json:"user_id"`
		ServiceID int32   `json:"service_id"`
		OrderID   int32   `json:"order_id"`
		Price     float32 `json:"price"`
		Comment   string  `json:"comment"`
	}
	ReportRequest struct {
		Year  string `json:"year"`
		Month string `json:"month"`
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeReceiptReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req ReceiptRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeReserveReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req ReserveRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeReportReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req ReportRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeGetBalanceReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetBalanceRequest
	vars := mux.Vars(r)

	req = GetBalanceRequest{
		Id: vars["id"],
	}
	return req, nil
}
