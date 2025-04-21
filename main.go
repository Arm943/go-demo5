package main

import (
	"flag"
	"fmt"
	"weather/geo"
	"weather/weather"
)

func main() {
	fmt.Println("Новый проект")
	city := flag.String("city", "", "Город пользователя")
	format := flag.Int("format", 1, "формат вывода погоды")

	flag.Parse()

	fmt.Println(*city)
	geoDate, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
	}

	weatherDate := weather.GetWeather(*geoDate, *format)
	fmt.Println(weatherDate)
}
