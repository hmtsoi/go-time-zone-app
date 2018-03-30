package rest

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"thmlogwork.com/go-time-zone-app/domain"
)

type TimezoneController struct {
	TimezoneService domain.TimezoneService
}

func (controller TimezoneController) Create() {

	router := mux.NewRouter()
	router.HandleFunc("/timeForLatLon/{latLonStr}", controller.GetTimezone).Methods("GET")

	log.Print("Start http router listeing to 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func (controller TimezoneController) GetTimezone(writer http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)
	latLonStr := params["latLonStr"]
	latLon, err := validateAndParseLatLonInput(latLonStr)

	if err != nil {
		log.Print(err)
		http.Error(writer, err.Error(), 400)
		return
	}

	log.Printf("Getting info for (Latitude, Longitude): (%v, %v)\n", latLon.Latitude, latLon.Longitude)
	info := controller.TimezoneService.GetTimezoneInfo(latLon)
	response := controller.TimezoneService.MapTimezoneInfo(info)
	json.NewEncoder(writer).Encode(response)
}

func validateAndParseLatLonInput(latLonStr string) (domain.LatLon, error) {
	arr := strings.Split(latLonStr, ",")
	if len(arr) != 2 {
		return domain.LatLon{}, errors.New(
			"Please input longitude and latitude comma separated in form of {latitude},{longitude}")
	}

	lat, err1 := strconv.ParseFloat(arr[0], 64)
	lon, err2 := strconv.ParseFloat(arr[1], 64)
	var err error
	if err1 != nil {
		err = err1
	} else {
		err = err2
	}
	return domain.LatLon{Latitude: lat, Longitude: lon}, err

}
