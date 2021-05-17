package gomedo

import (
	"log"
	"time"
)

func Watch() {
	log.Printf("watching %s with a %ds interval\n", C.UniqueID, int(C.Interval.Seconds()))

	var prevIDs map[uint64]interface{}

	for true {
		as, err := Scrape()

		if err != nil {
			log.Println(err)
			return
		}

		allIDs := make(map[uint64]interface{})

		var newAppointments []AppointmentResponse

		for _, a := range *as {
			allIDs[a.ID] = nil

			if _, ok := prevIDs[a.ID]; !ok {
				newAppointments = append(newAppointments, a)
			}
		}

		prevIDs = allIDs

		log.Printf("%d relevant new appoinments\n", len(newAppointments))

		if err := Notify(&newAppointments); err != nil {
			log.Println(err)
			return
		}

		time.Sleep(C.Interval)
	}
}
