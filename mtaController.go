package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	LAT     string
	LONG    string
	API_KEY string
}

var (
	env                 Env
	url_near_bus_routes *url.URL
	near_bus_stops      map[string]interface{}
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	env := Env{LAT: os.Getenv("LAT"), LONG: os.Getenv("LONG"), API_KEY: os.Getenv("MTA_API")}

	// copy paste from mta website
	mta, err := url.Parse("https://bustime.mta.info/api/where/stops-for-location.json?lat=40.748433&lon=-73.985656&latSpan=0.005&lonSpan=0.005&key=YOUR_KEY_HERE")
	if err != nil {
		log.Fatal(err)
	}

	q := mta.Query()
	q.Set("lat", env.LAT)
	q.Set("lon", env.LONG)
	q.Set("key", env.API_KEY)
	mta.RawQuery = q.Encode()

	url_near_bus_routes = mta
}

func getNearBuses() {
	response, err := http.Get(url_near_bus_routes.String())
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	json.Unmarshal(body, &near_bus_stops)
	fmt.Print(near_bus_stops)

}
