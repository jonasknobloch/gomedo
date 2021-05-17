package gomedo

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Response struct {
	Response     string                `json:"response"`
	Appointments []AppointmentResponse `json:"termine"`
}

type AppointmentResponse struct {
	Date         string `json:"date"`
	Time         string `json:"time"`
	ID           uint64 `json:"id"`
	Practitioner string `json:"practitioner"`
	Title        string `json:"title"`
}

func (a *AppointmentResponse) UnmarshalJSON(d []byte) error {
	var tmp []interface{}

	if err := json.Unmarshal(d, &tmp); err != nil {
		return err
	}

	if id, err := strconv.ParseUint(tmp[2].(string), 10, 64); err == nil {
		a.ID = id
	} else {
		return err
	}

	a.Date = tmp[0].(string)
	a.Time = tmp[1].(string)
	a.Practitioner = tmp[3].(string)
	a.Title = tmp[4].(string)

	return nil
}

func Scrape() (*[]AppointmentResponse, error) {
	log.Printf("scraping %s\n", C.Endpoint)

	data := url.Values{}
	data.Set("uniqueident", C.UniqueID)

	client := &http.Client{}
	req, err := http.NewRequest("POST", C.Endpoint, strings.NewReader(data.Encode()))

	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	r := Response{}

	if err := json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	if r.Response != "ok" {
		return nil, errors.New("response not okay")
	}

	log.Printf("scraped %d appointments\n", len(r.Appointments))

	return filterAppointmentResponses(&r.Appointments, C.Keywords), nil
}

func filterAppointmentResponses(as *[]AppointmentResponse, ks []string) *[]AppointmentResponse {
	if len(ks) == 0 {
		return as
	}

	var r []AppointmentResponse
	for _, a := range *as {
		l := strings.ToLower(a.Title)
		for _, k := range ks {
			if strings.Contains(l, k) {
				r = append(r, a)
			}
			break
		}
	}

	return &r
}
