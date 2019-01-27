package location

import (
	"errors"
	"net"

	"github.com/mesg-foundation/core/client/service"
)

const locateLocationOutputKey = "location"

type locateInputs struct {
	IP string `json:"ip"`
}

type locateLocationOutputs struct {
	City           string  `json:"city"`
	Country        string  `json:"country"`
	CountryISOCode string  `json:"countryISOCode"`
	Continent      string  `json:"continent"`
	ContinentCode  string  `json:"continentCode"`
	TimeZone       string  `json:"timeZone"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	PostalCode     string  `json:"postalCode"`
}

type locationData struct {
	City struct {
		Names struct {
			EN string `maxminddb:"en"`
		} `maxminddb:"names"`
	} `maxminddb:"city"`
	Continent struct {
		Code  string `maxminddb:"code"`
		Names struct {
			EN string `maxminddb:"en"`
		} `maxminddb:"names"`
	} `maxminddb:"continent"`
	Country struct {
		Code  string `maxminddb:"iso_code"`
		Names struct {
			EN string `maxminddb:"en"`
		} `maxminddb:"names"`
	} `maxminddb:"country"`
	Location struct {
		Latitude  float64 `maxminddb:"latitude"`
		Longitude float64 `maxminddb:"longitude"`
		TimeZone  string  `maxminddb:"time_zone"`
	} `maxminddb:"location"`
	Postal struct {
		Code string `maxminddb:"code"`
	} `maxminddb:"postal"`
}

// locateHandler is a task handler to locate an IP address.
func (s *LocationService) locateHandler(execution *service.Execution) (string, interface{}) {
	var inputs locateInputs
	if err := execution.Data(&inputs); err != nil {
		return newErrorOutput(err)
	}

	ip := net.ParseIP(inputs.IP)
	if ip == nil {
		return newErrorOutput(errors.New("IP address is not valid"))
	}

	var data locationData
	if err := s.db.Lookup(ip, &data); err != nil {
		return newErrorOutput(err)
	}

	return "location", locateLocationOutputs{
		City:           data.City.Names.EN,
		Country:        data.Country.Names.EN,
		CountryISOCode: data.Country.Code,
		Continent:      data.Continent.Names.EN,
		ContinentCode:  data.Continent.Code,
		TimeZone:       data.Location.TimeZone,
		Latitude:       data.Location.Latitude,
		Longitude:      data.Location.Longitude,
		PostalCode:     data.Postal.Code,
	}
}
