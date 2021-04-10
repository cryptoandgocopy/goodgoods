package data

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"goodgoods/utils"
)

type countryGoods struct {
	Country     string `json:"country"`
	Goods       string `json:"goods"`
	ChildLabor  string `json:"child_labor"`
	ForcedLabor string `json:"forced_labor"`
}

const yes = "Yes"

/*
IsGood requests and process data to check if
specified goods from specified origin are ethical
*/
func IsGood(origin string, goods string) bool {
	// request data
	list := parseCountryGoods(readFile(utils.Config.DB_Path))

	// check if goods match
	var status = true

	for _, item := range list {
		// first match search criteria (case insensitive)
		if (strings.EqualFold(item.Country, origin)) && (strings.EqualFold(item.Goods, goods)) {
			// now check if any unethical labor matches
			if (strings.EqualFold(item.ChildLabor, yes)) || (strings.EqualFold(item.ForcedLabor, yes)) {
				status = false
			}
		}
	}

	return status
}

/*
Open and read data file
*/
func readFile(path string) []byte {
	// open file
	jsonFile, err := os.Open(path)
	defer jsonFile.Close()
	utils.CheckErr(err)

	// read contents
	byteValue, err := ioutil.ReadAll(jsonFile)
	utils.CheckErr(err)

	return byteValue
}

/*
Parse JSON data into struct
*/
func parseCountryGoods(response []byte) []countryGoods {
	// read into struct
	var list []countryGoods

	err := json.Unmarshal(response, &list)
	utils.CheckErr(err)

	return list
}
