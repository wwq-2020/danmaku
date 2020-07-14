package server

import (
	"net"

	"github.com/sirupsen/logrus"
	"github.com/wwq-2020/danmaku/pkg/conf"
)

// Server Server
type Server struct {
	ln net.Listener
}

// MustNew MustNew
func MustNew(conf *conf.Conf) *Server {
	server, err := New(conf)
	if err != nil {
		logrus.WithField("err", err).
			Fatalf("failed to new server")
	}
	return server
}

// New New
func New(conf *conf.Conf) (*Server, error) {
	ln, err := net.Listen("tcp4", conf.Server.Addr)
	if err != nil {
		return nil, err
	}

	return &Server{ln: ln}, nil
}

// Start Start
func (s *Server) Start() error {
	return nil
}

// Stop Stop
func (s *Server) Stop() error {
	return nil
}
