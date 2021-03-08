package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	quasar "github.com/rafaelraba/quasar_fire/internal"
	"github.com/rafaelraba/quasar_fire/internal/plataform/server/handler/health"
	satellites "github.com/rafaelraba/quasar_fire/internal/plataform/server/handler/satelites"
	"log"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine
	dataInterpreter quasar.DataInterpreter
}

func New(host string, port string, interpreter quasar.DataInterpreter) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:"+ port, host),
		dataInterpreter: interpreter,
	}
	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/topsecret", satellites.TopsecretHandler(s.dataInterpreter))
	s.engine.GET("/topsecret", satellites.TopSecretGetHandler(s.dataInterpreter))
	s.engine.POST("/topsecret/:satellite_name", satellites.TopSecretSaveHandler(s.dataInterpreter))
}
