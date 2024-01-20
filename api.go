package main

import (
	"context"
	"encoding/json"
	"fmt"
	"gnana997/price-fetcher-go/types"
	"math/rand"
	"net/http"
)

type APIFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

type JSONAPIServer struct {
	listenAddr string
	svc        PriceFetcher
}

func NewJSONAPIServer(listenAddr string, svc PriceFetcher) *JSONAPIServer {
	return &JSONAPIServer{
		listenAddr: listenAddr,
		svc:        svc,
	}
}

func (s *JSONAPIServer) Run() {
	http.HandleFunc("/", makeHTTPAPIFunc(s.handleFetchPrice))
	http.ListenAndServe(s.listenAddr, nil)
}

func makeHTTPAPIFunc(apiFn APIFunc) http.HandlerFunc {
	fmt.Println("Inside makeHTTpAPIFunc")
	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestId", rand.Intn(1000000))
	return func(w http.ResponseWriter, r *http.Request) {
		if err := apiFn(ctx, w, r); err != nil {
			writeJson(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}
}

func (s *JSONAPIServer) handleFetchPrice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Inside handleFetchPrice")
	ticker := r.URL.Query().Get("ticker")
	price, err := s.svc.FetchPrice(ctx, ticker)
	if err != nil {
		return err
	}
	priceResp := types.PriceResponse{
		Ticker: ticker,
		Price:  price,
	}

	return writeJson(w, http.StatusOK, priceResp)
}

func writeJson(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)
	return json.NewEncoder(w).Encode(v)
}
