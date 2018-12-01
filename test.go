package main

import (
	"net/http"
	"log"
	"encoding/json"
	"io/ioutil"
)

func main() {

	res, _ := http.Get(`http://localhost:5000/graphql?query={ticker(symbol:"AAPL"){latest_price,company_name}}`)

	body, err := ioutil.ReadAll(res.Body)

	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	var jsonBlob = []byte(body)

	type TickerData struct {
		CompanyName string `json:"company_name"`
		LatestPrice string `json:"latest_price"`
	}

	type Ticker struct {
		TickerData TickerData `json:"ticker"`
	}

	type Data struct {
		Data  Ticker `json:"data"`
	}

	var data Data

	err = json.Unmarshal(jsonBlob, &data)

	log.Println(data)

}