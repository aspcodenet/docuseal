package jobs

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func Start() {
	c := cron.New()
	c.AddFunc("@hourly", func() { fmt.Println("Hourly job") })
	c.Start()
}
