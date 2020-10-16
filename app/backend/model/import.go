package model

import "time"

type Import struct {
	ImportId string
	Date time.Time
	Note string
	ProviderId string
}
