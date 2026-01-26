package main

import "fmt"

type Pessoa struct {
	Nome   string
	Idade  uint
	Peso   float64
	Altura float64
	Cpf    string
}

type Smartphone struct {
	Marca         string
	Armazenamento uint
	Cor           string
	Peso          float64
	Valor         float64
	Ligado        bool
	Proprietario  Pessoa
}

func (s *Smartphone) Ligar() {
	s.Ligado = true
}

func (s *Smartphone) Desligar() {
	s.Ligado = false
}

type IOS struct {
	Smartphone
	OS string
}

func (i *IOS) DownloadAppleStore() {
	fmt.Println("baixando....")
}

type Android struct {
	Smartphone
	OS string
}

func (a *Android) DownloadGooglePlay() {

	if a.Ligado {
		fmt.Println("Baixando google play...")
		return
	}

	fmt.Println("O celular está desligado!")
}

func main() {

	iphone := IOS{
		Smartphone: Smartphone{
			Marca:         "Apple",
			Armazenamento: 128,
			Cor:           "prata",
			Peso:          130,
			Valor:         10000,
			Ligado:        false,
		},
		OS: "ios X",
	}

	fmt.Println(iphone)

	fmt.Println("Ligando iphone...")
	iphone.Ligar()
	fmt.Println(iphone.Ligado)

	iphone.DownloadAppleStore()

	s20 := Android{
		Smartphone: Smartphone{
			Marca:         "Samsung",
			Armazenamento: 256,
			Cor:           "prata",
			Peso:          125,
			Valor:         7655,
		},
		OS: "Android x.xx",
	}

	s20.DownloadGooglePlay()
	s20.Ligar()
	s20.DownloadGooglePlay()

}
