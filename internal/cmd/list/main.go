package main

import (
	"fmt"
	"strconv"

	"tonek.org/tenn/internal/activecomm"
)

func main() {
	date := "2021-08-24"
	cl := activecomm.NewClient()
	courts, err := cl.FindResources("woodland")
	check(err)
	var periods []activecomm.DateTimePeriod
	for i := 8; i <= 22; i++ {
		from := fmt.Sprintf("%s %02d:00:00", date, i)
		to := fmt.Sprintf("%s %02d:59:00", date, i)
		periods = append(periods, activecomm.DateTimePeriod{
			From: from,
			To:   to,
		})
	}
	for _, c := range courts.Body.Items {
		res, err := cl.GetReservations(activecomm.ReservationTimeGroupRequest{
			ResourceID:      strconv.Itoa(int(c.ID)),
			Periods:         periods,
			ReservationUnit: 1,
		})
		check(err)
		fmt.Println(c.Name)
		for _, r := range res.Body.ReservationTimes {
			fmt.Printf("%v - %v: %v\n", r.Start, r.End, r.Availability)
		}
		fmt.Println()
	}
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
