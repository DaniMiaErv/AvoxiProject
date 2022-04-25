package checkip

type GetCheckIPRequest struct {
	IpAd        string   `json:"ipAd"`
	CountryList []string `json:"countryList,omitempty"`
}
