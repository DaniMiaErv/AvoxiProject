package checkip

type PassFailResponse struct {
	PassFail bool  `json:"passFail"`
	Error    error `json:"error"`
}
