package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var (
	env                 Env
	url_near_bus_routes *url.URL
	near_bus_stops      StopsForLocationType
	stops_on_lines      StopOnLinesType
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	env = Env{LAT: os.Getenv("LAT"), LONG: os.Getenv("LONG"), API_KEY: os.Getenv("MTA_API")}

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

func getNearBusStops() {
	response, err := http.Get(url_near_bus_routes.String())
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	json.Unmarshal(body, &near_bus_stops)

	getClosestBuses()
}

func getClosestBuses() {
	mta, err := url.Parse("https://bustime.mta.info/api/siri/stop-monitoring.json?key=&OperatorRef=MTA&MonitoringRef=&LineRef=&MaximumStopVisits=2&StopMonitoringDetailLevel=minimum")
	if err != nil {
		log.Fatal(err)
	}

	stops_on_lines.Code = near_bus_stops.Code
	stops_on_lines.CurrentTime = near_bus_stops.CurrentTime
	stops_on_lines.Data.LimitExceeded = near_bus_stops.Data.LimitExceeded

	for _, stop := range near_bus_stops.Data.Stops {
		newStop := StopOnLinesStops{Code: stop.Code, Direction: stop.Direction, Name: stop.Name}
		for _, route := range stop.Routes {
			newRoute := StopOnLinesRoute{Description: route.Description, ID: route.ID, LongName: route.LongName, ShortName: route.ShortName}
			var json_res StopMonitoringType
			q := mta.Query()
			q.Set("key", env.API_KEY)
			q.Set("MonitoringRef", stop.Code)
			q.Set("LineRef", route.ID)
			mta.RawQuery = q.Encode()
			response, err := http.Get(mta.String())
			if err != nil {
				log.Fatal(err)
			}
			defer response.Body.Close()
			body, err := io.ReadAll(response.Body)
			json.Unmarshal(body, &json_res)

			if smd := json_res.Siri.ServiceDelivery.StopMonitoringDelivery; len(smd) > 0 && len(smd[0].MonitoredStopVisit) > 0 {
				for _, bus := range smd[0].MonitoredStopVisit {
					newBus := Bus{bus.MonitoredVehicleJourney.ProgressRate, bus.MonitoredVehicleJourney.MonitoredCall.AimedArrivalTime, bus.MonitoredVehicleJourney.MonitoredCall.ExpectedArrivalTime, bus.MonitoredVehicleJourney.MonitoredCall.Extensions.Distances.PresentableDistance, bus.MonitoredVehicleJourney.MonitoredCall.Extensions.Distances.StopsFromCall}
					newRoute.Buses = append(newRoute.Buses, newBus)
				}
			}
			newStop.Routes = append(newStop.Routes, newRoute)
		}
		stops_on_lines.Data.Stops = append(stops_on_lines.Data.Stops, newStop)
	}
}

func compileBusFacts() []string {
	var rv []string
	if stops_on_lines.Code != 200 {
		return []string{"error: mta code != 200"}
	}
	if stops_on_lines.Data.LimitExceeded {
		return []string{"error: data limit exceeded"}
	}
	for _, stop := range stops_on_lines.Data.Stops {
		for _, route := range stop.Routes {
			if len(route.Buses) > 0 {
				pT1, _ := time.Parse("2006-01-02T15:04:05.000-07:00", route.Buses[0].AimedArrivalTime)
				if len(route.Buses) > 1 {
					pT2, _ := time.Parse("2006-01-02T15:04:05.000-07:00", route.Buses[1].AimedArrivalTime)
					rv = append(rv, fmt.Sprintf("%s %s -> %s %s @ %s S:%d & %s S:%d", stop.Name, stop.Direction, route.ShortName, route.LongName, pT1.Format("15:04"), route.Buses[0].StopsFromCall, pT2.Format("15:04"), route.Buses[1].StopsFromCall))
				} else {
					rv = append(rv, fmt.Sprintf("%s %s -> %s %s @ %s S:%d", stop.Name, stop.Direction, route.ShortName, route.LongName, pT1.Format("15:04"), route.Buses[0].StopsFromCall))
				}
			} else {
				rv = append(rv, fmt.Sprintf("%s %s -> %s %s @ no bus, sadness overwhelms me :(", stop.Name, stop.Direction, route.ShortName, route.LongName))
			}
		}
	}
	return rv
}
