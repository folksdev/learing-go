package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/leekchan/accounting"
)

type Time struct {
	Updated    string
	UpdatedIso string
	UpdatedUk  string
}

type Currency struct {
	Code       string
	Symbol     string
	Rate       string
	Rate_Float float64
}

type Bpi struct {
	Usd Currency
	Eur Currency
	Gbp Currency
}

type ApiResponse struct {
	Time      Time
	ChartName string
	Bpi       Bpi
}

func main() {
	response, err := http.Get("https://api.coindesk.com/v1/bpi/currentprice.json")
	if err != nil {
		log.Fatal(err)
	}

	if response.StatusCode == 200 {
		data, _ := ioutil.ReadAll(response.Body)
		//fmt.Println(string(data))

		var apiResponse ApiResponse
		json.Unmarshal(data, &apiResponse)

		//fmt.Println(apiResponse)

		ac := accounting.Accounting{Symbol: "$", Precision: 4}
		fmt.Println("Last Updated:", apiResponse.Time.Updated, " Bitcoin (USD): ", ac.FormatMoney(apiResponse.Bpi.Usd.Rate_Float))

	} else {
		fmt.Println("Http call failed with status: ", response.Status)
	}

}
