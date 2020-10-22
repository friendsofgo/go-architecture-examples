package http

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	ReadTimeoutDefault  = 10 // seconds
	WriteTimeoutDefault = 10 // seconds
)

type Server struct {
	*gin.Engine
	httpAddr string
}

type Handler interface {
	Router(r gin.IRouter)
}

func NewServer(httpAddr string) Server {
	return Server{
		Engine:   gin.New(),
		httpAddr: httpAddr,
	}
}

func (s Server) Run() error {
	httpSrv := &http.Server{
		Addr:         s.httpAddr,
		Handler:      s.Engine,
		ReadTimeout:  ReadTimeoutDefault * time.Second,
		WriteTimeout: WriteTimeoutDefault * time.Second,
	}

	return httpSrv.ListenAndServe()
}
