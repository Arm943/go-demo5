package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type GeoDate struct {
	City string `json:"city"`
}

func GetMyLocation(city string) (*GeoDate, error) {
	if city != "" {
		isCity := checkCity(city)
		if !isCity {
			panic("Такого города нет!")
		}
		return &GeoDate{
			City: city,
		}, nil
	}
	resp, err := http.Get("https://ipapi.co/json/")
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("NOT200")
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var geo GeoDate
	json.Unmarshal(body, &geo)
	return &geo, nil
}

type CityPopulationResponce struct {
	Error bool `json:"error"`
}

func checkCity(city string) bool {
	postBody, _ := json.Marshal(map[string]string{
		"city": city,
	})

	resp, err := http.Post("https://countriesnow.space/api/v0.1/countries/population/cities", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	var populationResponce CityPopulationResponce
	json.Unmarshal(body, &populationResponce)
	return !populationResponce.Error
}
