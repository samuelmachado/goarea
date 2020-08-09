package goarea

import "math"

// presenta o valor da razão entre a circunferência de qualquer círculo e seu diâmetro.
const Pi = 3.1416

// Circ é responsável por calcular a área do círculo
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

func Rect(base, altura float64) float64 {
	return base * altura
}

func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
