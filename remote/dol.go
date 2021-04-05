package remote

import (
	"encoding/json"
	"fmt"
	"strings"

	"goodgoods/utils"

	fasthttp "github.com/valyala/fasthttp"
)

type countryGoods struct {
	ID               string `json:"id"`
	AssessmentID     string `json:"assessment_id"`
	Year             string `json:"year"`
	CountryID        string `json:"country_id"`
	Country          string `json:"country"`
	GoodID           string `json:"good_id"`
	Good             string `json:"good"`
	RegionName       string `json:"regionname"`
	ChildLabor       string `json:"child_labor"`
	ForcedLabor      string `json:"forced_labor"`
	ForcedChildLabor string `json:"forced_child_labor"`
}

const yes = "Yes"

/*
IsGood requests and process remote data to check if
specified goods from specified origin are ethical
*/
func IsGood(origin string, goods string) (bool, error) {
	// request remote data
	response, err := fetch(utils.Config.DOL_API_URL_CountryGoods)
	utils.CheckErr(err)

	// parse remote data
	list, err := parseCountryGoods(response)
	utils.CheckErr(err)

	// check if goods match
	var status = true

	for _, item := range list {
		// first match search criteria (case insensitive)
		if (strings.EqualFold(item.Country, origin)) && (strings.EqualFold(item.Good, goods)) {
			// now check if any unethical labor matches
			if (strings.EqualFold(item.ChildLabor, yes)) || (strings.EqualFold(item.ForcedLabor, yes)) || (strings.EqualFold(item.ForcedChildLabor, yes)) {
				status = false
			}
		}
	}

	return status, err
}

/*
Parse remote JSON data into struct
*/
func parseCountryGoods(response string) ([]countryGoods, error) {
	// read into struct
	var list []countryGoods
	err := json.Unmarshal([]byte(response), &list)
	utils.CheckErr(err)

	return list, nil
}

/*
Handle remote calls to Department of Labor
*/
func fetch(url string) (string, error) {
	// build request
	request := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(request)
	request.SetRequestURI(url)
	request.Header.Set("X-API-KEY", utils.Config.DOL_API_KEY)

	// build response
	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)

	// execute
	err := fasthttp.Do(request, response)
	if err != nil {
		return "", fmt.Errorf("Client get failed: %s", err)
	}

	if response.StatusCode() != fasthttp.StatusOK {
		return "", fmt.Errorf("Expected status code %d but got %d", fasthttp.StatusOK, response.StatusCode())
	}

	return string(response.Body()), nil
}
