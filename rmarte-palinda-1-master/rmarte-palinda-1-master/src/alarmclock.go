package main

import (
	"fmt"
	"time"
)

func Remind(text string, delay time.Duration) {
	for {
		<-time.After(delay * time.Second)
		fmt.Println("The time is", time.Now().Format("15.04.05")+":", text)
	}
}

func main() {
	go Remind("Time to eat", 10)
	go Remind("Time to work", 30)
	go Remind("Time to sleep", 60)
	select {}
}
