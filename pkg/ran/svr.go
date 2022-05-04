package ran

import (
	"net/url"
)

type Server struct {
	Core
}

func NewServer(url *url.URL) (*Server, error) {
	svr := &Server{
		NewCore(),
	}
	return svr, nil
}

func (s *Server) Run() (stmtCh StmtChannel, err error) {
	go s.Core.RunCoreLoop()
	return s.stmtCh, nil
}
