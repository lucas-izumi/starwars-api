package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "strings"
    "strconv"
    "starwarsrest/get_api_info"
)

type Planet struct {
    ID        string   `json:"id,omitempty"`
    Name 	  string   `json:"name,omitempty"`
    Climate   string   `json:"climate,omitempty"`
    Terrain   string   `json:"terrain,omitempty"`
    Films     int      `json:"films,omitempty"`
}

var planets_ []Planet
var auto_id_ int

//Add planet (name, climate, terrain)
func AddPlanetEndpoint(w http.ResponseWriter, req *http.Request) {
    var planet Planet
    _ = json.NewDecoder(req.Body).Decode(&planet)
    auto_id_ = auto_id_ + 1
    planet.ID = strconv.FormatInt(int64(auto_id_), 10)
    planet.Films = GetNumElements(getapiinfo.GetApiInformation(planet.Name))
    planets_ = append(planets_, planet)
    json.NewEncoder(w).Encode(planets_)
}

//List planets
func GetPlanetsEndpoint(w http.ResponseWriter, req *http.Request) {
    json.NewEncoder(w).Encode(planets_)
}

//Search a planet either by ID or name
func SearchPlanetEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for _, item := range planets_ {
        if item.ID == params["search"] || strings.EqualFold(item.Name, params["search"]) {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Planet{})
}

//Delete a planet
func DeletePlanetEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    for index, item := range planets_ {
        if item.ID == params["id"] {
            planets_ = append(planets_[:index], planets_[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(planets_)
}

//Function to simply count elements of an array
func GetNumElements(results []string) int {
	i := 0
    for _, _ = range results {
			i++
    }
    return i
}

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/planets", GetPlanetsEndpoint).Methods("GET")
    router.HandleFunc("/planets/{search}", SearchPlanetEndpoint).Methods("GET")
    router.HandleFunc("/planets/", AddPlanetEndpoint).Methods("POST")
    router.HandleFunc("/planets/{id}", DeletePlanetEndpoint).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":12345", router))
}
