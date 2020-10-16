package model

import "time"

type Customer struct {
    CustomerId string
    CustomerName string
    DOB time.Time
    PhoneNumber string
    Address string
}
