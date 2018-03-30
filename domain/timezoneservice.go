package domain

type TimezoneService interface {
	GetTimezoneInfo(l LatLon) *TimezoneInfo
	MapTimezoneInfo(i *TimezoneInfo) *TimezoneInfoResponse
}
