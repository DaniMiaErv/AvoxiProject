package checkip

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

func GetCheckIPHandler(ep endpoint.Endpoint, options []httptransport.ServerOption) *httptransport.Server {
	return httptransport.NewServer(
		ep,
		decodeGetPassFailRequest,
		encodeGetPassFailResponse,
		options...,
	)

}

func decodeGetPassFailRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request GetCheckIPRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeGetPassFailResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	response, ok := response.(*PassFailResponse)
	if !ok {
		return errors.New("error decoding")
	}
	return json.NewEncoder(w).Encode(response)
}
