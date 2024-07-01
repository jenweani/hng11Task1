package services

import (
	"encoding/json"
	// "fmt"
	"io"
	"net/http"
	"os"
)


func GetTempByLoc(Ip string) (map[string]map[string]interface{}, error) {
	url := "https://weatherapi-com.p.rapidapi.com/current.json?q=" + Ip 

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", os.Getenv("RAPIDAPI_KEY"))
	req.Header.Add("x-rapidapi-host", "weatherapi-com.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return map[string]map[string]interface{}{}, err
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	
	if err != nil {
		return map[string]map[string]interface{}{}, err
	}

	var temp map[string]map[string]interface{}
	err = json.Unmarshal(body, &temp)
	if err != nil {
		return map[string]map[string]interface{}{}, err
	}

	
	return temp, nil
}