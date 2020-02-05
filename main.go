package main

import "time"

func main() {
	for {
		Handler()
		time.Sleep(time.Second * 60 * 30)
	}
}
