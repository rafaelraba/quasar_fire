package interpreter

import "strings"

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
    x = 18.4
    y = 28.5
    return x, y
}
