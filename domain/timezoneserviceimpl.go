package domain

import (
	"time"
)

type TimezoneServiceImpl struct {
	Repository TimezoneRepository
}

func (service *TimezoneServiceImpl) GetTimezoneInfo(l LatLon) *TimezoneInfo {
	return service.Repository.Get(l)
}

func (service *TimezoneServiceImpl) MapTimezoneInfo(i *TimezoneInfo) *TimezoneInfoResponse {
	loc, _ := time.LoadLocation(i.TimezoneParameter)
	now := time.Now()
	localTime := now.In(loc)
	utcTime := now.UTC()
	return &TimezoneInfoResponse{
		i.Places, i.DstPlaces, i.UtcFormat, i.TimezoneParameter,
		localTime, utcTime}
}
