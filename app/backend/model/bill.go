package model

import "time"

type Bill struct {
	BillId uint32
	Date time.Time
	CustomerId string
	Note string
}
