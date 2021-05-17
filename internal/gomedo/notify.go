package gomedo

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func Notify(as *[]AppointmentResponse) error {
	if len(*as) == 0 {
		log.Println("nothing to notify about")
		return nil
	}

	body, err := json.Marshal(*as)

	if err != nil {
		return err
	}

	for _, h := range C.Hooks {
		if err := triggerHook(h, body); err != nil {
			return err
		}
	}

	return nil
}

func triggerHook(h string, b []byte) error {
	req, err := http.NewRequest("POST", h, bytes.NewBuffer(b))

	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New(res.Status)
	}

	log.Printf("triggered hook %s\n", h)
	return nil
}
