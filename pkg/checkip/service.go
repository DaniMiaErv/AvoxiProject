package checkip

import (
	"context" //in case of concurrent requests
)

type Service interface {
	Get(ct context.Context, ipAd string, countryList []string) (bool, error)
	ServiceStatus(ct context.Context) (int, error)
}
