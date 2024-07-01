package services

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type LocDetails struct {
	City string
	Country string
	Ip string
	Latitude float32
	Longitude float32
	Time_zone string
	Time_zone_offset string
  }

func GetLocationByIP(ip string) (LocDetails, error) {
	
	url := "https://ip-reputation-geoip-and-detect-vpn.p.rapidapi.com/?ip=" + ip

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", os.Getenv("RAPIDAPI_KEY"))
	req.Header.Add("x-rapidapi-host", "ip-reputation-geoip-and-detect-vpn.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return LocDetails{}, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocDetails{}, err
	}
	var locDetails LocDetails
	err = json.Unmarshal(body, &locDetails)
	if err != nil {
		return LocDetails{}, err
	}

	return locDetails, nil

}