package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)


func GetTempByLoc(lat, long float32) (interface{}, error) {
	url := "https://weatherapi-com.p.rapidapi.com/current.json?q=" + fmt.Sprintf("%.2f",lat) + "%2C"  + fmt.Sprintf("%.2f",long) 

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", os.Getenv("RAPIDAPI_KEY"))
	req.Header.Add("x-rapidapi-host", "weatherapi-com.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	
	if err != nil {
		return "", err
	}

	var temp map[string]map[string]interface{}
	err = json.Unmarshal(body, &temp)
	if err != nil {
		return "", err
	}
    t := temp["current"]["temp_c"]

	
	return fmt.Sprintf("%.2f", t), nil
}