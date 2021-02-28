package bootstrap

import "github.com/rafaelraba/quasar_fire/internal/plataform/server"

const (
    host = "localhost"
    port = 8080
)

func Run() error {
    srv := server.New(host,port)
    return srv.Run()
}