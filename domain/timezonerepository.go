package domain

type TimezoneRepository interface {
	Get(l LatLon) *TimezoneInfo
}
