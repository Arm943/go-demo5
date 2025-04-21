package geo

import (
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
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var geo GeoDate
	json.Unmarshal(body, &geo)
	if err != nil {
		return nil, err
	}
	return &geo, nil
}
