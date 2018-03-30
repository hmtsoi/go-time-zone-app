package domain

import "time"

type TimezoneInfoResponse struct {
	Places            string
	DstPlaces         string
	UtcFormat         string
	TimezoneParameter string
	CurrentLocalTime  time.Time
	CurrentUtcTime    time.Time
}
