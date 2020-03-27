package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "strings"
)

type Planet struct {
    ID        string   `json:"id,omitempty"`
    Name 	  string   `json:"name,omitempty"`
    Climate   string   `json:"climate,omitempty"`
    Terrain   string   `json:"terrain,omitempty"`
    Films     string   `json:"films,omitempty"`
}

var planets_ []Planet

//Add planet (name, climate, terrain)
func AddPlanetEndpoint(w http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    var planet Planet
    _ = json.NewDecoder(req.Body).Decode(&planet)
    planet.ID = params["id"]
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

func main() {
    router := mux.NewRouter()
    router.HandleFunc("/planets", GetPlanetsEndpoint).Methods("GET")
    router.HandleFunc("/planets/{search}", SearchPlanetEndpoint).Methods("GET")
    router.HandleFunc("/add-planet/{id}", AddPlanetEndpoint).Methods("POST")
    log.Fatal(http.ListenAndServe(":12345", router))
}
