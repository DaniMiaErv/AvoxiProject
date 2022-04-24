package endpoints

type GetRequest struct {
	ipAd        string
	countryList []string
}

type GetResponse struct {
	passFail bool
	Err      string
}

type ServiceStatusRequest struct{}

type ServiceStatusResponse struct {
	Code int
	Err  string
}
