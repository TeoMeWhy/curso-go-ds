package main

import (
	"fmt"
	"time"
)

type Personagem struct {
	Nome   string
	Raca   string
	Classe string
	Nivel  int
	Forca  int
	Defesa int
	Vida   int
}

type Jogador struct {
	Personagem
	TempoGameplay int64
	UltimoLogin   time.Time
	Online        bool
}

func (j *Jogador) CalcForcaAtaque() int {
	return j.Forca + j.Nivel + 10
}

func (j *Jogador) CalcForcaDefesa() int {
	return j.Defesa + j.Nivel
}

type Criatura struct {
	Personagem
	RespawTempo time.Duration
}

func (c *Criatura) CalcForcaAtaque() int {
	return c.Forca
}

func (c *Criatura) CalcForcaDefesa() int {
	return c.Defesa
}

type NPC struct {
	Personagem
	Funcao string
}

type Atacante interface {
	CalcForcaAtaque() int
}

type Defensor interface {
	CalcForcaDefesa() int
}

type Lutador interface {
	CalcForcaAtaque() int
	CalcForcaDefesa() int
}

func CalculoDano(a Atacante, d Defensor) int {
	return a.CalcForcaAtaque() - d.CalcForcaDefesa()
}

func RinhaDeGo(p1, p2 Lutador) {
	fmt.Println(CalculoDano(p1, p2))
	fmt.Println(CalculoDano(p2, p1))
}

func main() {

	teo := Jogador{
		Personagem: Personagem{
			Nome:   "teocalvo",
			Raca:   "humano",
			Classe: "mago",
			Nivel:  1,
			Forca:  2,
			Defesa: 0,
			Vida:   10,
		},
		TempoGameplay: 60,
		UltimoLogin:   time.Date(2025, 12, 4, 3, 10, 3, 0, &time.Location{}),
		Online:        true,
	}

	orc := Criatura{
		Personagem: Personagem{
			Nome:   "orc 1",
			Raca:   "orc",
			Classe: "guerreiro",
			Nivel:  2,
			Forca:  3,
			Defesa: 1,
			Vida:   5,
		},
		RespawTempo: 10,
	}

	// wardo := NPC{
	// 	Personagem: Personagem{
	// 		Nome:   "Wardo",
	// 		Raca:   "goblin",
	// 		Classe: "artesao",
	// 		Nivel:  0,
	// 		Forca:  1,
	// 		Defesa: 0,
	// 		Vida:   1,
	// 	},
	// }

	fmt.Println(teo)
	fmt.Println(orc)

	fmt.Println("Força de ataque Teo:", teo.CalcForcaAtaque())
	fmt.Println("Força de ataque Orc:", orc.CalcForcaAtaque())

	fmt.Println("Dano Téo:", CalculoDano(&teo, &orc))
	fmt.Println("Dano Orc:", CalculoDano(&orc, &teo))

	RinhaDeGo(&teo, &orc)

}
