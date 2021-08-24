package main

import (
	"fmt"

	"tonek.org/tenn/internal/activecomm"
)

func main() {
	cl := activecomm.NewClient()

	resp, err := cl.GetReservations(activecomm.ReservationTimeGroupRequest{
		ResourceID: "355",
		Periods: []activecomm.DateTimePeriod{
			{From: "2021-08-24 08:00:00", To: "2021-08-24 08:59:00"},
			{From: "2021-08-24 09:00:00", To: "2021-08-24 09:59:00"},
			{From: "2021-08-24 11:00:00", To: "2021-08-24 11:59:00"},
		},
		ReservationUnit: 1,
	})
	check(err)
	fmt.Printf("%+v\n", resp)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
