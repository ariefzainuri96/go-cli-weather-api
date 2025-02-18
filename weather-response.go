// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    weatherResponse, err := UnmarshalWeatherResponse(bytes)
//    bytes, err = weatherResponse.Marshal()

package main

import "encoding/json"

func UnmarshalWeatherResponse(data []byte) (WeatherResponse, error) {
	var r WeatherResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *WeatherResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type WeatherResponse struct {
	QueryCost         json.Number         `json:"queryCost"`
	Latitude          float64             `json:"latitude"`
	Longitude         float64             `json:"longitude"`
	ResolvedAddress   string              `json:"resolvedAddress"`
	Address           string              `json:"address"`
	Days              []CurrentConditions `json:"days"`
	CurrentConditions CurrentConditions   `json:"currentConditions"`
}

type CurrentConditions struct {
	Datetime       string              `json:"datetime"`
	DatetimeEpoch  json.Number         `json:"datetimeEpoch"`
	Temp           float64             `json:"temp"`
	Feelslike      float64             `json:"feelslike"`
	Humidity       float64             `json:"humidity"`
	Dew            float64             `json:"dew"`
	Precip         float64             `json:"precip"`
	Precipprob     float64             `json:"precipprob"`
	Snow           json.Number         `json:"snow"`
	Snowdepth      json.Number         `json:"snowdepth"`
	Preciptype     []string            `json:"preciptype"`
	Windgust       float64             `json:"windgust"`
	Windspeed      float64             `json:"windspeed"`
	Winddir        float64             `json:"winddir"`
	Pressure       float64             `json:"pressure"`
	Visibility     float64             `json:"visibility"`
	Cloudcover     float64             `json:"cloudcover"`
	Solarradiation float64             `json:"solarradiation"`
	Solarenergy    float64             `json:"solarenergy"`
	Uvindex        json.Number         `json:"uvindex"`
	Severerisk     json.Number         `json:"severerisk"`
	Conditions     string              `json:"conditions"`
	Icon           string              `json:"icon"`
	Stations       []string            `json:"stations"`
	Source         string              `json:"source"`
	Sunrise        *string             `json:"sunrise,omitempty"`
	SunriseEpoch   *json.Number        `json:"sunriseEpoch,omitempty"`
	Sunset         *string             `json:"sunset,omitempty"`
	SunsetEpoch    *json.Number        `json:"sunsetEpoch,omitempty"`
	Moonphase      *float64            `json:"moonphase,omitempty"`
	Tempmax        *float64            `json:"tempmax,omitempty"`
	Tempmin        *float64            `json:"tempmin,omitempty"`
	Feelslikemax   *float64            `json:"feelslikemax,omitempty"`
	Feelslikemin   *float64            `json:"feelslikemin,omitempty"`
	Precipcover    *float64            `json:"precipcover,omitempty"`
	Description    *string             `json:"description,omitempty"`
	Hours          []CurrentConditions `json:"hours,omitempty"`
}
