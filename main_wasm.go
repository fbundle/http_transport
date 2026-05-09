//go:build js && wasm

package main

import (
	"encoding/json"
	"net/http"

	ht "github.com/fbundle/http_transport/http_transport"
)

func main() {
	r := ht.New()

	r.POST("api/echo", func(body []byte) (int, any) {
		var req map[string]any
		if err := json.Unmarshal(body, &req); err != nil {
			return http.StatusBadRequest, nil
		}
		return http.StatusOK, req
	})

	select {}
}
