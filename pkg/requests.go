package checkip

type GetPassFailRequest struct {
	ipAd        string   `json:"ipAd"`
	countryList []string `json:"countryList"`
}

// type Set struct {
// 	GetEndpoint           endpoint.Endpoint
// 	ServiceStatusEndpoint endpoint.Endpoint
// }

// func NewEndpointSet(svc checkip.Service) Set {
// 	return Set{
// 		GetEndpoint:           MakeGetEndpoint(svc),
// 		ServiceStatusEndpoint: MakeServiceStatusEndpoint(svc),
// 	}
// }

// func MakeGetEndpoint(svc checkip.Service) endpoint.Endpoint {
// 	return func(ct context.Context, request interface{}) (interface{}, error) {
// 		req := request.(GetRequest)
// 		passResponse, err := svc.Get(ct, req.ipAd, req.countryList)
// 		if err != nil {
// 			return GetResponse{passResponse, err.Error()}, nil
// 		}

// 		return GetResponse{passResponse, ""}, nil
// 	}
// }

// func MakeServiceStatusEndpoint(svc checkip.Service) endpoint.Endpoint{
// 	return func(ct context.Context, request interface{}) (interface{}, error){
// 		_ = request.(ServiceStatusRequest)
// 		code, err := svc.ServiceStatus(ct)
// 		if err != nil{
// 			return ServiceStatusResponse{Code: code, Err: ""}, nil
// 		}
// 		return ServiceStatusResponse{Code: code, Err: ""}, nil
// 	}
// }

// func (s *Set) Get(ct context.Context, ipAd string, countryList []string)(bool, error){
// 	response, err := s.GetEndpoint(ct, GetRequest{ipAd: string, countryList: []string})
// 	if err != nil{
// 		return passFail, err
// 	}

// 	getResponse := response.(GetResponse)
// 	if getResponse.Err != ""{
// 		return passFail, errors.New(getResponse.Err)
// 	}
// 	return getResponse.passFail, nil
// }

// func (s *Set) ServiceStatus(ct context.Context) (int, error) {
// 	response, err := s.ServiceStatusEndpoint(ct, ServiceStatusRequest{})
// 	svcStatusResp := response.(ServiceStatusResponse)
// 	if err != nil {
// 		return svcStatusResp.Code, err
// 	}
// 	if svcStatusResp.Err != "" {
// 		return svcStatusResp.Code, errors.New(svcStatusResp.Err)
// 	}
// 	return svcStatusResp.Code, nil
// }

// var logger log.Logger

// func init(){
// 	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
// 	logger = log.With(logger, "timestamp", log.DefaultTimestampUTC)
// }
