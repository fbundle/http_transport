//go:build !js

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	ht "github.com/fbundle/http_transport/http_transport"
)

func main() {
	mux := http.NewServeMux()
	r := ht.NewGo(mux)

	r.POST("api/echo", func(body []byte) (int, any) {
		var req map[string]any
		if err := json.Unmarshal(body, &req); err != nil {
			return http.StatusBadRequest, nil
		}
		return http.StatusOK, req
	})

	mux.Handle("/", http.FileServer(http.Dir("docs")))
	fmt.Println("Listening on http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
