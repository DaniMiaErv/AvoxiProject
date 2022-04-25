package checkip

import (
	"context" //in case of concurrent requests
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-kit/kit/log"
)

type Service interface {
	GetIPCheck(ctx context.Context, ipAd string, countryList []string) (bool, error)
	//ServiceStatus(ctx context.Context) (int, error)
}

type checkipService struct {
	log log.Logger
}

func NewService(logger log.Logger) Service {
	return &checkipService{
		log: logger,
	}
}

func (svc *checkipService) GetIPCheck(ctx context.Context, ipAd string, countryList []string) (bool, error) {
	//func NewService() Service { return &checkipService{} }

	//Step 1. Find what country ip address is from.
	//Step 2. Put the list of countries into a map.
	//Step 4. Compare it to the whitelisted countries. (contains)

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
