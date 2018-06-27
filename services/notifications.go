package services

import (
	"fmt"
	"time"
)

func Pending(action string) *time.Ticker {
	fmt.Print(action)
	ticker := time.NewTicker(time.Second)
	go func() {
		for range ticker.C {
			fmt.Print(".")
		}
	}()
	return ticker
}
