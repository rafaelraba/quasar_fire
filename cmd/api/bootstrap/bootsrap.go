package bootstrap

import (
    "github.com/rafaelraba/quasar_fire/internal/plataform/interpreter"
    "github.com/rafaelraba/quasar_fire/internal/plataform/server"
    "log"
    "os"
)

const (
    host = "localhost"
    port = "8080"
)

func Run() error {
    dataInterpreter := interpreter.MainInterpreter{}
    var srv server.Server
    portEnv := os.Getenv("PORT")
    if  portEnv != ""  {
        log.Println("DEBUG ENV", os.Getenv("PORT"))
        srv = server.New("", portEnv ,dataInterpreter)
        return srv.Run()
    }
    srv = server.New(host, port, dataInterpreter)
    return srv.Run()
}
