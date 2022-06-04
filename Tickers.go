package main

//This exercise is related to channels and will be developed in the future

//Timers are for when you want to do something once in the future
// tickers are for when you want to do something repeatedly at
//regular intervals.
//Here’s an example of a ticker that ticks periodically until we stop it.

import (
	"fmt"
	"time"
)

//Tickers use a similar mechanism to timers: a channel that is sent values.
//Here we’ll use the select builtin on the channel to await the values as they arrive every 500ms.

func main() {

	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	//Tickers can be stopped like timers. Once a ticker is stopped
	// it won’t receive any more values on its channel.
	//We’ll stop ours after 1600ms.

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")

}

//When we run this program the ticker should tick 3 times before we stop it.

//output :

//Tick at 2022-06-04 11:29:56.487625 -0700 PDT
//Tick at 2022-06-04 11:29:56.988063 -0700 PDT
//Tick at 2022-06-04 11:29:57.488076 -0700 PDT
//Ticker stopped
