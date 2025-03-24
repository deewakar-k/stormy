package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	api_key      = "your-api-key"
	default_city = "Hyderabad" //TODO: geo location
	units        = "metric"    //default
	timeplus     = 0
	timeminus    = 0

	// ANSI color codes
	yellow    = "\033[38;5;227m"
	lightBlue = "\033[38;5;153m"
	purple    = "\033[38;5;147m"
	reset     = "\033[0m"
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
	city := default_city
	if len(os.Args) > 1 {
		city = os.Args[1]
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&units=%s&APPID=%s", city, units, api_key)

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

	// better way to do this?
	sourcesunrise := int64(weather.Sys.Sunrise)
	sourcesunset := int64(weather.Sys.Sunset)

	sunriseDatetime := time.Unix(sourcesunrise, 0).Local()
	sunsetDatetime := time.Unix(sourcesunset, 0).Local()

	adjustedSunrise := sunriseDatetime.Add(time.Duration(timeplus) * time.Hour).Add(-time.Duration(timeminus) * time.Hour)
	adjustedSunset := sunsetDatetime.Add(time.Duration(timeplus) * time.Hour).Add(-time.Duration(timeminus) * time.Hour)

	sunrisestring := adjustedSunrise.Format(time.Kitchen)
	sunsetstring := adjustedSunset.Format(time.Kitchen)

	var weather_condition string
	var output []string

	weather_condition = strings.ToLower(weather.Weather[0].Main)

	if weather_condition == "clear" {
		output = []string{
			"                 " + "󰖙  weather: clear",
			"      " + yellow + ".-." + reset + "      " + "󰔏  temperature: " + temp,
			"   " + yellow + "‒ (   ) ‒" + reset + "   " + "󰖝  wind speed: " + wind_speed,
			"      " + yellow + "`-᾿" + reset + "      " + "󰖜  sunrise: " + sunrisestring,
			"     " + yellow + "/   \\" + reset + "     " + "󰖛  sunset: " + sunsetstring,
		}
	} else if weather_condition == "clouds" {
		output = []string{
			"                 " + "󰖐  weather: cloudy",
			"       .--.      " + "󰔏  temperature: " + temp,
			"    .-(    ).    " + "󰖝  wind speed: " + wind_speed,
			"   (___.__)__)   " + "󰖜  sunrise: " + sunrisestring,
			"                 " + "󰖛  sunset: " + sunsetstring,
		}
	} else if weather_condition == "rain" {
		output = []string{
			"                 " + "󰖗  weather: rainy",
			"       .--.      " + "󰔏  temperature: " + temp,
			"    .-(    ).    " + "󰖝  wind speed: " + wind_speed,
			"   (___.__)__)   " + "󰖜  sunrise: " + sunrisestring,
			"    " + lightBlue + "ʻ‚ʻ‚ʻ‚ʻ‚ʻ" + reset + "    " + "󰖛  sunset: " + sunsetstring,
		}
	} else if weather_condition == "snow" {
		output = []string{
			"                 " + "󰖘  weather: snowy",
			"    .-(    ).    " + "󰔏  temperature: " + temp,
			"   (___.__)__)   " + "󰖝  wind speed: " + wind_speed,
			"     * * * *     " + "󰖜  sunrise: " + sunrisestring,
			"    * * * *      " + "󰖛  sunset: " + sunsetstring,
		}
	} else if weather_condition == "thunderstorm" {
		output = []string{
			"                 " + "󰖓  weather: stormy",
			"    .-(    ).    " + "󰔏  temperature: " + temp,
			"   (___.__)__)   " + "󰖝  wind speed: " + wind_speed,
			"      " + purple + "  /_" + reset + "       " + "󰖜  sunrise: " + sunrisestring,
			"       " + purple + "  /" + reset + "       " + "󰖛  sunset: " + sunsetstring,
		}
	} else if weather_condition == "haze" {
		output = []string{
			"                 " + "󰖑  weather: hazy",
			"    ~ ~ ~ ~      " + "󰔏  temperature: " + temp,
			"   ~ ~ ~ ~ ~     " + "󰖝  wind speed: " + wind_speed,
			"    ~ ~ ~ ~      " + "󰖜  sunrise: " + sunrisestring,
			"                 " + "󰖛  sunset: " + sunsetstring,
		}
	} else {
		output = []string{
			"                 " + "󰖐  weather: " + weather_condition,
			"    .-(    ).    " + "󰔏  temperature: " + temp,
			"   (___.__)__)   " + "󰖝  wind speed: " + wind_speed,
			"                 " + "󰖜  sunrise: " + sunrisestring,
			"                 " + "󰖛  sunset: " + sunsetstring,
		}
	}

	fmt.Println(strings.Join(output, "\n"))
}
