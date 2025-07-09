package main

import (
	"net/http"
	"fmt"
)

type ProxyHandler struct{}

func (h *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    fmt.Fprintf(w, "Proxy server is running!")
}