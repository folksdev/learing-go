package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/leekchan/accounting"
)

type Time struct {
	Updated    string
	UpdatedIso string
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
}

type ApiResponse struct {
	Time      Time
	ChartName string
	Bpi       Bpi
}

func main() {
	fmt.Println("Starting the application...")
	response, err := http.Get("https://api.coindesk.com/v1/bpi/currentprice.json")
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		//fmt.Println(string(data))
		var parsed ApiResponse
		json.Unmarshal(data, &parsed)

		//fmt.Println(parsed)

		ac := accounting.Accounting{Symbol: "$", Precision: 4}
		fmt.Println(parsed.Time.Updated, " - Bitcoin (USD)", ac.FormatMoney(parsed.Bpi.Usd.Rate_Float))
	}

}
