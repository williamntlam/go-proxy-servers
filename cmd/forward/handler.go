package main

import (
	"net/http"
	"fmt"
)

func (h *ProxyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// 1. Extract target URL from request
    // 2. Forward request to target server  
    // 3. Return response back to client

    fmt.Fprintf(w, "Proxy server is running!")
}