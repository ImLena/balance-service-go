package balance

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("POST").Path("/receipt").Handler(httptransport.NewServer(
		endpoints.Receipt,
		decodeReceiptReq,
		encodeResponse,
	))

	r.Methods("POST").Path("/reserve").Handler(httptransport.NewServer(
		endpoints.Reserve,
		decodeReserveReq,
		encodeResponse,
	))

	r.Methods("POST").Path("/accept").Handler(httptransport.NewServer(
		endpoints.AcceptPayment,
		decodeReserveReq,
		encodeResponse,
	))

	r.Methods("GET").Path("/{id}/balance").Handler(httptransport.NewServer(
		endpoints.GetBalance,
		decodeGetBalanceReq,
		encodeResponse,
	))

	r.Methods("POST").Path("/report").Handler(httptransport.NewServer(
		endpoints.Report,
		decodeReportReq,
		encodeResponse,
	))

	return r

}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
