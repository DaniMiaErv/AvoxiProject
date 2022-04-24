package checkip

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
)

type checkipService struct{}

//constructor
func NewService() Service { return &checkipService{} }

func (c *checkipService) Get(_ context.Context, ipAd string, countryList []string) (bool, error) {

	//api gateway

	// func HandleRequest(request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error){

	// }

	// func CheckIPWhiteList(ipAd string, countryList []string) bool {
	//func CheckIPHandler(w http.ResponseWriter, r *http.Request){
	//Step 0. Check that valid IP address was passed.
	//Step 1. Find what country ip address is from.
	//Step 2. Put the list of countries into a map.
	//Step 3. Check if the list passed is empty.
	//Step 4. Compare it to the whitelisted countries. (contains)

	//create a get request to return fields:
	//status, message (if status fails), country and query)

	response, err := http.Get("http://ip-api.com/json/" + ipAd + "?fields=status,message,country,query")
	if err != nil {
		fmt.Printf("Response to check IP failed: %s\n", err)
	} else {

		//close after use to prevent possible leaks
		defer response.Body.Close()

		var data map[string]string

		json.NewDecoder(response.Body).Decode(&data)
		if err != nil {
			fmt.Printf("Error unmarshaling response: %s\n", err)
		}

		//check that status of api request passed.
		if data["message"] != "nil" {
			fmt.Printf("Status failed: %v\n ", data["message"])
		}

		//get the ip address country and convert
		//to json string
		//ipCountry, err := json.Marshal(data["country"])

		if err != nil {
			fmt.Printf("Error marshaling country IP Ad comes from: %s\n", err)
		}

		//195 possible countries in the world.
		//put given whitelisted countries into map

		//create empty map
		countries := make(map[string]int)

		//copy contents into the map
		for index, element := range countryList {
			countries[element] = index
		}

		//check if ip address comes from
		//whitelisted country
		_, exists := countries[data["country"]]

		return exists, nil
	}

	return false, nil

}

func (c *checkipService) ServiceStatus(_ context.Context) (int, error) {
	log.Println("Checking status of the ip check service")
	return http.StatusOK, nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "timestamp", log.DefaultTimestampUTC)
}
