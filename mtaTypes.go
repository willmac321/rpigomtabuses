package main

type Env struct {
	LAT     string
	LONG    string
	API_KEY string
}

type Bus struct {
	ProgressRate        string `json:"ProgressRate"`
	AimedArrivalTime    string `json:"AimedArrivalTime"`
	ExpectedArrivalTime string `json:"ExpectedArrivalTime"`
	PresentableDistance string `json:"PresentableDistance"`
	StopsFromCall       int    `json:"StopsFromCall"`
}

type StopOnLinesRoute struct {
	Description string `json:"description"`
	ID          string `json:"id"`
	LongName    string `json:"longName"`
	ShortName   string `json:"shortName"`
	Buses       []Bus  `json:"buses"`
}

type StopOnLinesStops struct {
	Code      string             `json:"code"`
	Direction string             `json:"direction"`
	Name      string             `json:"name"`
	Routes    []StopOnLinesRoute `json:"routes"`
}

type StopOnLinesType struct {
	Code        int   `json:"code"`
	CurrentTime int64 `json:"currentTime"`
	Data        struct {
		LimitExceeded bool               `json:"limitExceeded"`
		Stops         []StopOnLinesStops `json:"stops"`
	} `json:"data"`
}

// https://mholt.github.io/json-to-go/
type StopsForLocationType struct {
	Code        int   `json:"code"`
	CurrentTime int64 `json:"currentTime"`
	Data        struct {
		LimitExceeded bool `json:"limitExceeded"`
		Stops         []struct {
			Code         string  `json:"code"`
			Direction    string  `json:"direction"`
			ID           string  `json:"id"`
			Lat          float64 `json:"lat"`
			LocationType int     `json:"locationType"`
			Lon          float64 `json:"lon"`
			Name         string  `json:"name"`
			Routes       []struct {
				Agency struct {
					Disclaimer     string `json:"disclaimer"`
					Email          string `json:"email"`
					FareURL        string `json:"fareUrl"`
					ID             string `json:"id"`
					Lang           string `json:"lang"`
					Name           string `json:"name"`
					Phone          string `json:"phone"`
					PrivateService bool   `json:"privateService"`
					Timezone       string `json:"timezone"`
					URL            string `json:"url"`
				} `json:"agency"`
				Color       string `json:"color"`
				Description string `json:"description"`
				ID          string `json:"id"`
				LongName    string `json:"longName"`
				ShortName   string `json:"shortName"`
				TextColor   string `json:"textColor"`
				Type        int    `json:"type"`
				URL         string `json:"url"`
			} `json:"routes"`
			WheelchairBoarding string `json:"wheelchairBoarding"`
		} `json:"stops"`
	} `json:"data"`
	Text    string `json:"text"`
	Version int    `json:"version"`
}

type StopMonitoringType struct {
	Siri struct {
		ServiceDelivery struct {
			ResponseTimestamp      string `json:"ResponseTimestamp,omitempty"`
			StopMonitoringDelivery []struct {
				ResponseTimestamp  string `json:"ResponseTimestamp,omitempty"`
				ValidUntil         string `json:"ValidUntil,omitempty"`
				MonitoredStopVisit []struct {
					RecordedAtTime          string `json:"RecordedAtTime,omitempty"`
					MonitoredVehicleJourney struct {
						LineRef                 string `json:"LineRef,omitempty"`
						DirectionRef            string `json:"DirectionRef,omitempty"`
						FramedVehicleJourneyRef struct {
							DataFrameRef           string `json:"DataFrameRef,omitempty"`
							DatedVehicleJourneyRef string `json:"DatedVehicleJourneyRef,omitempty"`
						} `json:"FramedVehicleJourneyRef,omitempty"`
						JourneyPatternRef        string `json:"JourneyPatternRef,omitempty"`
						PublishedLineName        string `json:"PublishedLineName,omitempty"`
						OperatorRef              string `json:"OperatorRef,omitempty"`
						OriginRef                string `json:"OriginRef,omitempty"`
						DestinationRef           string `json:"DestinationRef,omitempty"`
						DestinationName          string `json:"DestinationName,omitempty"`
						OriginAimedDepartureTime string `json:"OriginAimedDepartureTime,omitempty"`
						SituationRef             []struct {
							SituationSimpleRef string `json:"SituationSimpleRef,omitempty"`
						} `json:"SituationRef,omitempty"`
						Monitored       bool `json:"Monitored,omitempty"`
						VehicleLocation struct {
							Longitude float64 `json:"Longitude,omitempty"`
							Latitude  float64 `json:"Latitude,omitempty"`
						} `json:"VehicleLocation,omitempty"`
						Bearing        float64 `json:"Bearing,omitempty"`
						ProgressRate   string  `json:"ProgressRate,omitempty"`
						ProgressStatus string  `json:"ProgressStatus,omitempty"`
						Occupancy      string  `json:"Occupancy,omitempty"`
						VehicleRef     string  `json:"VehicleRef,omitempty"`
						BlockRef       string  `json:"BlockRef,omitempty"`
						MonitoredCall  struct {
							StopPointRef          string `json:"StopPointRef,omitempty"`
							VisitNumber           int    `json:"VisitNumber,omitempty"`
							ExpectedArrivalTime   string `json:"ExpectedArrivalTime,omitempty"`
							ExpectedDepartureTime string `json:"ExpectedDepartureTime,omitempty"`
							Extensions            struct {
								Distances struct {
									PresentableDistance    string  `json:"PresentableDistance,omitempty"`
									DistanceFromCall       float64 `json:"DistanceFromCall,omitempty"`
									StopsFromCall          int     `json:"StopsFromCall,omitempty"`
									CallDistanceAlongRoute float64 `json:"CallDistanceAlongRoute,omitempty"`
								} `json:"Distances,omitempty"`
								Capacities struct {
									EstimatedPassengerCount    int `json:"EstimatedPassengerCount,omitempty"`
									EstimatedPassengerCapacity int `json:"EstimatedPassengerCapacity,omitempty"`
								} `json:"Capacities,omitempty"`
								VehicleFeatures struct {
									StrollerVehicle bool `json:"StrollerVehicle,omitempty"`
								} `json:"VehicleFeatures,omitempty"`
							} `json:"Extensions,omitempty"`
							StopPointName      string `json:"StopPointName,omitempty"`
							AimedArrivalTime   string `json:"AimedArrivalTime,omitempty"`
							AimedDepartureTime string `json:"AimedDepartureTime,omitempty"`
						} `json:"MonitoredCall,omitempty"`
						OnwardCalls struct{} `json:"OnwardCalls,omitempty"`
					} `json:"MonitoredVehicleJourney,omitempty,omitempty"`
				} `json:"MonitoredStopVisit,omitempty"`
			} `json:"StopMonitoringDelivery,omitempty"`
			SituationExchangeDelivery []struct {
				Situations struct {
					PtSituationElement []struct {
						PublicationWindow struct {
							StartTime string `json:"StartTime,omitempty"`
							EndTime   string `json:"EndTime,omitempty"`
						} `json:"PublicationWindow,omitempty"`
						Severity    string `json:"Severity,omitempty"`
						Summary     string `json:"Summary,omitempty"`
						Description string `json:"Description,omitempty"`
						Affects     struct {
							VehicleJourneys struct {
								AffectedVehicleJourney []struct {
									LineRef      string `json:"LineRef,omitempty"`
									DirectionRef string `json:"DirectionRef,omitempty"`
								} `json:"AffectedVehicleJourney,omitempty"`
							} `json:"VehicleJourneys,omitempty"`
						} `json:"Affects,omitempty"`
						CreationTime    string `json:"CreationTime,omitempty"`
						SituationNumber string `json:"SituationNumber,omitempty"`
					} `json:"PtSituationElement,omitempty"`
				} `json:"Situations,omitempty"`
			} `json:"SituationExchangeDelivery,omitempty"`
		} `json:"ServiceDelivery,omitempty"`
	} `json:"Siri,omitempty"`
}
