package server

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "log"
)

type Server struct {
    httpAddr string
    engine   *gin.Engine
}

func New(host string, port uint) Server {
    srv := Server{
        engine:   gin.New(),
        httpAddr: fmt.Sprintf("%s:%d", host, port),
    }
    return srv
}

func (s *Server) Run() error {
    log.Println("Server running on", s.httpAddr)
    return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
    s.engine.POST("")
}

