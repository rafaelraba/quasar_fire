package interpreter

import (
    satellites "github.com/rafaelraba/quasar_fire/internal/plataform/server/handler/satelites"
    "strings"
)

type MainInterpreter struct {
}

func (MainInterpreter) GetMessage(messages ...[]string) (msg string) {
    size := len(messages[0])
    if !validateSize(size, messages) {
        return "Tamaño diferente"
    }
    finalMessage := make([]string, size)
    for _, message := range messages {
        for index, word := range message {
            if word != "" {
                finalMessage[index] = word
            }
        }
    }
    return resolveMessage(finalMessage)
}

func validateSize(size int, messages [][]string) bool {
    for _, message := range messages {
        if len(message) != size {
            return false
        }
    }
    return true
}

func resolveMessage(finalMessage []string) string {
    for _, word := range finalMessage {
        if word == "" {
            return "No se pude determinar el mensaje"
        }
    }
    return strings.Join(finalMessage, " ")
}

func (MainInterpreter) GetLocation(distances ...float32) (x, y float32) {
    P1 := satellites.Position{
        X: -500,
        Y: -200,
    }
    rA := distances[0]

    P2 := satellites.Position{
        X: 100,
        Y: -100,
    }
    rB := distances[1]

    P3 := satellites.Position{
        X: 500,
        Y: 100,
    }
    rC := distances[2]

    ex := DivisionByScalar(Substract(P2, P1), VectorNorm(Substract(P2, P1)))
    i := DotProduct(ex, Substract(P3, P1))
    ey := DivisionByScalar(
        Substract(
            Substract(P3, P1),
            MultiplyByScalar(ex, i),
        ),
        VectorNorm(
            Substract(
                Substract(P3, P1),
                MultiplyByScalar(ex, i),
            ),
        ),
    )
    j := DotProduct(ey, Substract(P3, P1))

    d := VectorNorm(Substract(P2, P1))

    x = (Pow(rA, 2) - Pow(rB, 2) + Pow(d, 2)) / (2 * d)
    y = ((Pow(rA, 2) - Pow(rC, 2) + Pow(i, 2) + Pow(j, 2)) / (2 * j)) - ((i / j) * x)

    return float32(x), float32(y)
}
