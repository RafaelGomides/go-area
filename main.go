package goarea

import "math"

// Pi é uma proporção numérica definida pela relação entre o perímetro de uma circunferencia e seu diametro
const Pi = 3.1416

// Circ Calcula a area da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect é reposnsável por calcular a area do retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
