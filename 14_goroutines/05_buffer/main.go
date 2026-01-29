package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func pinger(c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 20; i++ {
		c <- "ping!"
		log.Println("Mensagem enviada! (ping)")
	}
	fmt.Println("Envio de Pinger Finalizado!!")
}

func ponger(c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 20; i++ {
		c <- "pong!"
		log.Println("Mensagem enviada! (pong)")
	}
	fmt.Println("Envio de Ponger Finalizado!!")
}

func printer(c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range c {
		log.Println("Mensagem recebida!", msg)
		time.Sleep(time.Millisecond * 250)
	}
	log.Println("Todas mensagens retiradas do canal!")
}

func main() {

	wgSenders := sync.WaitGroup{}
	canal := make(chan string, 40) // tem buffer de 40

	wgPrinter := sync.WaitGroup{}
	wgPrinter.Add(1)
	go printer(canal, &wgPrinter)

	wgSenders.Add(1)
	go pinger(canal, &wgSenders)

	wgSenders.Add(1)
	go ponger(canal, &wgSenders)

	wgSenders.Wait()
	close(canal)

	wgPrinter.Wait()

}
