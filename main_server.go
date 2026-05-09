//go:build !js

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	ht "github.com/fbundle/http_transport/http_transport"
)

func main() {
	port := os.Args[1]

	mux := http.NewServeMux()
	r := ht.NewGo(mux)

	r.POST("api/echo", func(body []byte) (int, []byte) {
		var req map[string]any
		if err := json.Unmarshal(body, &req); err != nil {
			return http.StatusBadRequest, nil
		}
		resp, _ := json.Marshal(req)
		return http.StatusOK, resp
	})

	mux.Handle("/", http.FileServer(http.Dir("docs")))
	fmt.Printf("Listening on http://localhost:%s\n", port)
	http.ListenAndServe(":"+port, mux)
}
