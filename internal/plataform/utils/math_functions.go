package utils

import (
    satellites "github.com/rafaelraba/quasar_fire/internal/plataform/server/handler/satelites"
    "math"
)

func Pow(base float32, exponent float32) float32 {
    return float32(math.Pow(float64(base), float64(exponent)))
}

// VectorNorm retorna un valor escalar equivalente a la norma de un vector
func VectorNorm(a satellites.Position) float32 {
    xPow := Pow(a.X, 2)
    yPow := Pow(a.Y, 2)
    norm := math.Sqrt(float64(xPow + yPow))
    return float32(norm)
}

// DotProduct retorna un valor escalar del producto punto de 2 coordenadas 2D
func DotProduct(a,b satellites.Position) float32 {
    return float32((a.X * b.X) + (a.Y * b.Y))
}

// Substract retorna una coordenada equivalente al resultado de: a - b
func Substract(a,b satellites.Position) satellites.Position {
    return satellites.Position{
        X: a.X - b.X,
        Y: a.Y - b.X,
    }
}

// DivisionByScalar retorna un vector dividido por un escalar
func DivisionByScalar(a satellites.Position, scalar float32) satellites.Position {
    //TODO: validar divison por 0
    return satellites.Position{
        X: a.X / scalar,
        Y: a.Y / scalar,
    }
}

// MultiplyByScalar retorna un vector multiplicado por un escalar
func MultiplyByScalar(a satellites.Position, scalar float32) satellites.Position {
    return satellites.Position{
        X: a.X * scalar,
        Y: a.Y * scalar,
    }
}