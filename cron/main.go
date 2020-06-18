package main

import (
	"fmt"

	cron "gopkg.in/robfig/cron.v2"
)

func main() {
	end := make(chan bool)
	c := cron.New()
	_, err := c.AddFunc("* * * * *", func() { fmt.Println("This runs every 1 min") })
	if err != nil {
		fmt.Println("There was an error:" + err.Error())
	}
	c.Start()
	<-end
}
