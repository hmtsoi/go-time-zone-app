package domain

type TimezoneServiceImpl struct {
	Repository TimezoneRepository
}

func (service *TimezoneServiceImpl) GetTimezoneInfo(l LatLon) *TimezoneInfoResponse {
	return service.Repository.Get(l)
}
