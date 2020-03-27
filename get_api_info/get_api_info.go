package getapiinfo

import (
    "encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type ApiResponse struct {
	Count int `json:"count"`
	Results []Results `json:"results"`
}

type Results struct {
	Films []string `json:"films"`
}

var httpConf = http.Client{Timeout: 10 * time.Second}

func GetApiInformation(pname string) []string {
	client := &http.Client{}

    req, _ := http.NewRequest("GET", "https://swapi.co/api/planets/", nil)
    req.Header.Add("Accept", "application/json")

    q := req.URL.Query()
    q.Add("search", pname)
    req.URL.RawQuery = q.Encode()

    resp, err := client.Do(req)

    if err != nil {
        log.Fatal("Errored when sending request to the server")
    }

    defer resp.Body.Close()
    resp_body, _ := ioutil.ReadAll(resp.Body)


    res := ApiResponse{}
	jsonErr := json.Unmarshal([]byte(resp_body), &res)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return res.Results[0].Films
}
