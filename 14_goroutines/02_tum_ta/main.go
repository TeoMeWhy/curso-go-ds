package main

import (
	"fmt"
	"log"
	"time"
)

func toque(txt string, t int) {
	for {
		log.Println(txt)
		time.Sleep(time.Second * time.Duration(t))
	}
}

func main() {

	go toque("Tum", 1)
	go toque("Tá", 2)

	stop := ""
	fmt.Scanf("%s", &stop)

}
