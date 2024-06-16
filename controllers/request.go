package controllers

import (
	"encoding/json"
	"fmt"
	"go-wa-bot-v3/structs"
	"io/ioutil"
	"math"
	"net/http"
	"os"

	"go-wa-bot-v3/utils"
)


func maxArray(array []float64) float64 {
	if len(array) == 0 {
		return 0
	}
	maxValue := array[0]
	for _, value := range array {
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

func GetData() (structs.Parameter, error){
	endpoint := os.Getenv("ENDPOINT")
	date := utils.GetYesterdayDate()
	url := fmt.Sprintf("%s/%s", endpoint, date)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching data: %v\n", err)
		return structs.Parameter{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return structs.Parameter{}, err
	}

	var response structs.Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return structs.Parameter{}, err
	}

	// Calculate the total sum of avg and max irradiance
	var totalAvg, totalMax float64
	var arrayRampRate, arrayMaxBeban []float64

	for i := 0; i < len(response.Data.Avg); i++ {
		totalAvg += response.Data.Avg[i].Irradiance / 360
		if i < len(response.Data.Max) {
			totalMax += response.Data.Max[i].Irradiance / 360
		}
		if i < len(response.Data.Avg)-1 {
			selisihRampRate := response.Data.Max[i+1].Irradiance - response.Data.Avg[i].Irradiance
			if !math.IsNaN(selisihRampRate) {
				arrayRampRate = append(arrayRampRate, selisihRampRate)
			}

			selisihMaxBeban := response.Data.Max[i].Irradiance - response.Data.Avg[i].Irradiance
			arrayMaxBeban = append(arrayMaxBeban, selisihMaxBeban)
		}
	}

	const conversionFactor = 6.8 * 0.1917

	maxRampRate := maxArray(arrayRampRate)
	rampRate := maxRampRate * conversionFactor
	if rampRate > 200 {
		rampRate = 200
	}

	maxBeban := maxArray(arrayMaxBeban)
	maxBebanBSS := maxBeban * conversionFactor
	if maxBebanBSS > 600 {
		maxBebanBSS = 600
	}

	forecastProduksiPV := totalAvg * conversionFactor
	forecastProduksiPVBSS := totalMax * conversionFactor
	forecastSmoothing := forecastProduksiPVBSS - forecastProduksiPV
	kebutuhanDoD := forecastSmoothing / 900 * 100

	parameters := structs.Parameter{
		ForecastProduksiPV : forecastProduksiPV,
		ForecastProduksiPVBSS : forecastProduksiPVBSS,
		ForecastSmoothing : forecastSmoothing,
		KebutuhanDoD : kebutuhanDoD,
		RampRate : rampRate,
		MaxBebanBSS : maxBebanBSS,
	}

	return parameters, nil
}