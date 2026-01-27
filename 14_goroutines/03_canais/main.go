package main

import (
	"fmt"
	"log"
	"time"
)

func ping(c chan string) {
	for {
		c <- "ping!"
		time.Sleep(time.Second * time.Duration(1))
	}
}

func pong(c chan string) {
	for {
		c <- "pong!"
		time.Sleep(time.Second * time.Duration(1))
	}
}

func print(c chan string) {
	for {
		msg := <-c
		log.Println(msg)
	}
}

func main() {

	canal := make(chan string)

	go ping(canal)
	go pong(canal)
	go print(canal)

	var ok string
	fmt.Scanf("%s", &ok)

}
