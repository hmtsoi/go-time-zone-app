package domain

type TimezoneService interface {
	GetTimezoneInfo(l LatLon) *TimezoneInfoResponse
}
