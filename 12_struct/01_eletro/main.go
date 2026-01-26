package main

import "fmt"

type Pessoa struct {
	Nome   string
	Idade  uint
	Peso   float64
	Altura float64
	Cpf    string
}

func (p Pessoa) GetIMC() float64 {
	return p.Peso / (p.Altura * p.Altura)
}

func (p *Pessoa) AddIdadeOne() {
	p.Idade++
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

func main() {

	teo := Pessoa{
		Nome:   "Téo Calvo",
		Idade:  34,
		Peso:   70.98,
		Altura: 1.8,
		Cpf:    "123.234.345-00",
	}

	celTeo := Smartphone{
		Marca:         "Motorola",
		Armazenamento: 128,
		Cor:           "Azul Piscina",
		Peso:          150.32,
		Valor:         1200.98,
		Proprietario:  teo,
	}

	fmt.Println(celTeo)

	fmt.Println("Marca:", celTeo.Marca)
	fmt.Println("Armazenamento:", celTeo.Armazenamento)
	celTeo.Valor += 250
	fmt.Println("Valor:", celTeo.Valor)
	fmt.Println("Ligado:", celTeo.Ligado)

	fmt.Println("Ligando celular....")
	celTeo.Ligar()
	fmt.Println("Ligado:", celTeo.Ligado)

	fmt.Println("Desligando celular....")
	celTeo.Desligar()
	fmt.Println("Ligado:", celTeo.Ligado)

	fmt.Println("IMC do proprietário:", celTeo.Proprietario.GetIMC())

	celTeo.Proprietario.AddIdadeOne()
	fmt.Println(celTeo)

	celNah := Smartphone{}
	fmt.Println(celNah)

}
