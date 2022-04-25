package checkip

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

type Endpoints struct {
	GetIPCheck endpoint.Endpoint
}

func MakeEndpoints(svc Service, logger log.Logger, middleware []endpoint.Middleware) Endpoints {
	return Endpoints{
		GetIPCheck: wrapEndpoint(makeCheckIPEndpoint(svc, logger), middleware),
	}
}

func makeCheckIPEndpoint(svc Service, logger log.Logger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req, ok := request.(*GetCheckIPRequest)
		if !ok {
			level.Error(logger).Log("message", "invalid request")
			return nil, errors.New("invalid request given")
		}

		passFail, err := svc.GetIPCheck(ctx, req.IpAd, req.CountryList)
		return &PassFailResponse{
			PassFail: passFail,
			Error:    err,
		}, nil
	}
}

func wrapEndpoint(endp endpoint.Endpoint, middleware []endpoint.Middleware) endpoint.Endpoint {
	for _, m := range middleware {
		endp = m(endp)
	}
	return endp
}
