package gomedo

import (
	"log"
	"os"
	"strings"
	"time"
)

type Config struct {
	UniqueID string
	Endpoint string
	Interval time.Duration
	Keywords []string
	Hooks    []string
}

var C = &Config{}

func init()  {
	C.UniqueID = os.Getenv("UNIQUE_IDENTIFIER")
	C.Endpoint = os.Getenv("SCRAPE_ENDPOINT")
	C.Interval, _ = time.ParseDuration(os.Getenv("SCRAPE_INTERVAL"))
	C.Keywords = strings.Split(os.Getenv("APPOINTMENT_KEYWORDS"), ",")
	C.Hooks = strings.Split(strings.TrimSpace(os.Getenv("NOTIFICATION_HOOKS")), ",")

	if !C.Valid() {
		panic("invalid configuration")
	}
}

func (c *Config) Valid() bool {
	valid := true

	valid = valid && !(c.UniqueID == "")
	valid = valid && !(c.Endpoint == "")
	valid = valid && !(c.Interval == 0 * time.Second)

	if len(c.Keywords) == 0 {
		log.Println("no appointment keywords set")
	}

	if len(c.Hooks) == 0 {
		log.Println("no notification hooks set")
	}

	return valid
}
