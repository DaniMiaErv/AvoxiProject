package main

//CheckIp "github.com/DaniMiaErv/AvoxiProject"
import (
	"net/http"
	"os"

	//checkip "github.com/DaniMiaErv/AvoxiProject/pkg/ipcheck"
	checkip "github.com/DaniMiaErv/AvoxiProject/pkg/checkip"

	"github.com/go-kit/kit/endpoint"
	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func main() {

	//testIp := "1"
	//countryList := []string{"United States", "Germany", "Japan"}

	//println(CheckIp.CheckIPWhiteList(testIp, countryList))

	var logger kitlog.Logger
	{
		logger = kitlog.NewLogfmtLogger(os.Stderr)
		logger = kitlog.With(logger, "timestamp", kitlog.DefaultTimestampUTC)
		logger = kitlog.With(logger, "caller", kitlog.DefaultCaller)
	}

	// var srv checkip.Service{
	// 	srv = checkip.
	// }
	var middlewares []endpoint.Middleware
	var options []httptransport.ServerOption
	svc := checkip.NewService(logger)
	eps := checkip.MakeEndpoints(svc, logger, middlewares)
	r := mux.NewRouter()
	r.Methods(http.MethodGet).Path("/checkip").Handler(checkip.GetCheckIPHandler(eps.GetIPCheck, options))

	level.Info(logger).Log("status", "listening", "port", "8080")
	svr := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: r,
	}
	level.Error(logger).Log(svr.ListenAndServe())
}
