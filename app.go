package main

import "fmt"

const ebulicaoC float64 = 100.0

//Sistema de conversao de ebulicao de Celsius para Kelvin
func main() {
	ebulicaoK := ebulicaoC + 273.0

	fmt.Printf("O ponto de ebulição da água em Kelvin é: %.2f K\n", ebulicaoK)
	fmt.Printf("O ponto de ebulição da água em Celsius é: %.2f °C\n", ebulicaoC)

}
