package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

const (
	city      = "Hyderabad" //TODO: system location
	units     = "metric"    //default
	timeplus  = 0
	timeminus = 0
)

type WeatherResponse struct {
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
	Sys struct {
		Sunrise int `json:"sunrise"`
		Sunset  int `json:"sunset"`
	} `json:"sys"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading .env")
	}

	appid := os.Getenv("WEATHERAPI_KEY")
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&units=%s&APPID=%s", city, units, appid)

	res, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error fetching weather data: %v\n", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "weather api status: %d\n", res.StatusCode)
		os.Exit(1)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading weather data: %v\n", err)
		os.Exit(1)
	}

	var weather WeatherResponse
	if err := json.Unmarshal(body, &weather); err != nil {
		fmt.Fprintf(os.Stderr, "error parsing weather data: %v\n", err)
		os.Exit(1)
	}

	var windspeedunits string

	if units == "metric" {
		windspeedunits = "m/s"
	} else if units == "imperial" {
		windspeedunits = "mph"
	} else {
		windspeedunits = "m/s"
	}

	wind_speed := fmt.Sprintf("%.2f %s", math.Round(weather.Wind.Speed), windspeedunits)

	temp := fmt.Sprintf("%.0f°", math.Round(weather.Main.Temp))

	sourcesunrise := int64(weather.Sys.Sunrise)
	sourcesunset := int64(weather.Sys.Sunset)

	sunriseDatetime := time.Unix(sourcesunrise, 0).UTC()
	sunsetDatetime := time.Unix(sourcesunset, 0).UTC()

	adjustedSunrise := sunriseDatetime.Add(time.Duration(timeplus) * time.Hour).Add(-time.Duration(timeminus) * time.Hour)
	adjustedSunset := sunsetDatetime.Add(time.Duration(timeplus) * time.Hour).Add(-time.Duration(timeminus) * time.Hour)

	sunrisestring := adjustedSunrise.Format("15:04:05")
	sunsetstring := adjustedSunset.Format("15:04:05")

	var weather_condition string
	var output []string

	weather_condition = strings.ToLower(weather.Weather[0].Main)

	if weather_condition == "clear" {
		output = []string{
			"     \\   /     " + "weather: clear",
			"      .-.      " + "temperature: " + temp,
			"   ‒ (   ) ‒   " + "wind speed: " + wind_speed,
			"      `-᾿      " + "sunrise: " + sunrisestring,
			"     /   \\     " + "sunset: " + sunsetstring,
		}
	} else if weather_condition == "clouds" {
		output = []string{
			"                 " + "weather: cloudy",
			"       .--.      " + "temprature: " + temp,
			"    .-(    ).    " + "wind speed: " + wind_speed,
			"   (___.__)__)   " + "sunrise: " + sunrisestring,
			"                 " + "sunset: " + sunsetstring,
		}
	} else if weather_condition == "rain" {
		output = []string{
			"                 " + "weather: rainy",
			"       .--.      " + "temperature: " + temp,
			"    .-(    ).    " + "wind speed: " + wind_speed,
			"   (___.__)__)   " + "sunrise: " + sunrisestring,
			"    ʻ‚ʻ‚ʻ‚ʻ‚ʻ    " + "sunset: " + sunsetstring,
		}
	} else if weather_condition == "snow" {
		output = []string{
			"                 " + "weather: snowy",
			"       .--.      " + "temperature: " + temp,
			"    .-(    ).    " + "wind speed: " + wind_speed,
			"   (___.__)__)   " + "sunrise: " + sunrisestring,
			"    ʻ‚ʻ‚ʻ‚ʻ‚ʻ    " + "sunset: " + sunsetstring,
		}
	} else if weather_condition == "thunderstorm" {
		output = []string{
			"       .--.      " + "weather: stormy",
			"    .-(    ).    " + "temperature: " + temp,
			"   (___.__)__)   " + "wind speed: " + wind_speed,
			"        /_       " + "sunrise: " + sunrisestring,
			"         /       " + "sunset: " + sunsetstring,
		}
	} else {
		output = []string{
			"       .--.      " + "weather: " + weather_condition,
			"    .-(    ).    " + "temperature: " + temp,
			"   (___.__)__)   " + "wind speed: " + wind_speed,
			"                 " + "sunrise: " + sunrisestring,
			"                 " + "sunset: " + sunsetstring,
		}
	}

	fmt.Println(strings.Join(output, "\n"))
}
