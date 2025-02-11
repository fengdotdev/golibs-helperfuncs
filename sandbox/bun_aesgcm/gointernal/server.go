package gointernal

import (
	"fmt"
	"net/http"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})


	http.HandleFunc("/encoded", ENCODED)

	if !Assert4charsAndNumbeable(port) {
		panic("Invalid port")
	}

	fmt.Println("Server running on port " + port)
	http.ListenAndServe(":"+port, nil)
}

func Assert4charsAndNumbeable(s string) bool {
	l := len(s)
	if !(l == 4) {
		return false
	}

	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}

	return true
}
