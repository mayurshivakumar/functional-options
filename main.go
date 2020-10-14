package main

import (
	"fmt"
	"github.com/pkg/errors"
)

type option func(server *Server) error

func Host(host string) option {
	return func(s *Server) error {
		// error check
		if host == "" {
			return errors.New("invalid host")
		}
		s.Host = host
		return nil
	}

}

func Port(port int) option {
	return func(s *Server) error {
		// dummy check for error.
		if (port % 2) == 0 {
			return errors.New("port in use")
		}
		s.Port = port
		return nil
	}
}

type Server struct {
	Host string
	Port int
}

func NewServer(opts ...option) (*Server, []error) {
	s := &Server{}
	errs := make([]error, 0)

	for _, opt := range opts {
		err := opt(s)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return s, errs
}

func main() {
	server1, errs := NewServer(Host("127.0.0.1"), Port(8001))
	if len(errs) > 0 {
		fmt.Println("invalid server1 config:", errs)
	} else {
		fmt.Printf("server1 host: %s, port: %d\n", server1.Host, server1.Port)
	}

	server2, errs := NewServer(Host("127.0.0.1"), Port(8000))
	if len(errs) > 0 {
		fmt.Println("invalid server2 config:", errs)
	} else {
		fmt.Printf("server2 host: %s, port: %d", server2.Host, server2.Port)
	}

}
