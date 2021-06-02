package main

import (
	"net/http"
	"os"

	"github.com/bluele/gcache"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	router *mux.Router
	log    *logrus.Logger
	port   string
	cache  gcache.Cache
}

func CreateServer() Server {
	s := Server{}
	err := s.Bootstrap()
	if err != nil {
		panic(err)
	}

	return s
}

func (s *Server) Bootstrap() error {
	// initialize logger
	l := logrus.New()
	s.log = l

	// set port number
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8888"
	}
	s.port = port

	// bind routes
	s.router = mux.NewRouter()
	s.router.HandleFunc("/", s.handleVersionWithFilter("stable"))
	s.router.HandleFunc("/all", s.handleVersionWithFilter("all"))
	s.router.HandleFunc("/alpha", s.handleVersionWithFilter("alpha"))
	s.router.HandleFunc("/rc", s.handleVersionWithFilter("rc"))

	// set up cache storage
	s.cache = gcache.New(20).LRU().Build()

	return nil
}

func (s *Server) Serve() error {
	s.log.Info("Starting server on port ", s.port)
	return http.ListenAndServe(":"+s.port, s.router)
}
