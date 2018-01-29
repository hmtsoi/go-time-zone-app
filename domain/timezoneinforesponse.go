package domain

type TimezoneInfoResponse struct {
	Places            string
	DstPlaces         string
	UtcFormat         string
	TimezoneParameter string
}

func NewTimezoneInfoResponse(
	places string,
	dstPlaces string,
	utcFormat string,
	timezoneParameter string) *TimezoneInfoResponse {

	return &TimezoneInfoResponse{
		Places:            places,
		DstPlaces:         dstPlaces,
		UtcFormat:         utcFormat,
		TimezoneParameter: timezoneParameter}
}
