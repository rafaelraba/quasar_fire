package bootstrap

import (
    "github.com/rafaelraba/quasar_fire/internal/plataform/interpreter"
    "github.com/rafaelraba/quasar_fire/internal/plataform/server"
)

const (
    host = "localhost"
    port = 8080
)

func Run() error {
    dataInterpreter := interpreter.MainInterpreter{}
    srv := server.New(host, port, dataInterpreter)
    return srv.Run()
}
