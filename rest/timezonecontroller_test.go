package rest

import (
	"testing"
	"thmlogwork.com/go-time-zone-app/domain"
)

func TestValidateAndParseLatLonInputWithCorrectInput(t *testing.T) {
	str := "18,2.3"
	output, _ := validateAndParseLatLonInput(str)
	latLon := domain.LatLon{Latitude: 18, Longitude: 2.3}
	if output != latLon {
		t.Errorf("Expected %v, but got %v", latLon, output)
	}
}

func TestValidateAndParseLatLonInputWithIncorrectInput(t *testing.T) {
	str := "18,2.3"
	_, err := validateAndParseLatLonInput(str)
	if err != nil {
		t.Error("Expected error")
	}
}

func TestValidateAndParseLatLonInputWithIncorrectNumber(t *testing.T) {
	str := "18,2e"
	_, err := validateAndParseLatLonInput(str)
	expectedError := "strconv.ParseFloat: parsing \"2e\": invalid syntax"
	if err.Error() != expectedError {
		t.Errorf("Expected \"%v\", but got \"%v\"", err, expectedError)
	}
}
