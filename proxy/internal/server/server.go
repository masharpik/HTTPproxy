package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/masharpik/HTTPproxy/internal"
)

func Init(host string, port string, handler internal.Handler) error {
	http.HandleFunc("/", handler.Proxy)

	addr := fmt.Sprintf("%s:%s", host, port)
	log.Printf("Server started on %s\n", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		err = fmt.Errorf("http.ListenAndServe: %w", err)
		return err
	}

	return nil
}
