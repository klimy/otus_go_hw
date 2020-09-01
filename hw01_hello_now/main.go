package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	localTime := time.Now()
	remoteTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("current time:", localTime.Round(0))
	fmt.Println("exact time:", remoteTime.Round(0))
}
