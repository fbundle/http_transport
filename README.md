# HTTP_TRANSPORT

## EXAMPLES

**Remote mode**
```bash
go run main_server.go 3000
# open http://localhost:3000
```

**Local (WASM) mode**
```bash
GOOS=js GOARCH=wasm go build -o docs/app.wasm main_wasm.go
cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" docs/
cd docs && python3 -m http.server 3000
# open http://localhost:3000
```

## PURPOSE

Write your backend handlers once in Go and run them in two modes without changing application code:

- **Remote** — handlers run as a standard HTTP server (`net/http`), the JS client calls them over the network via `fetch`
- **Local** — handlers compile to WebAssembly and run directly in the browser; the JS client calls them as plain functions with no network round-trip

The JS transport (`docs/http_transport.js`) auto-detects which mode is active at runtime.

## SPECIFICATION

- `http_transport.Router` — common interface (`POST(path, HandlerFunc)`)
- `http_transport.NewGo(mux)` — stdlib `net/http` implementation
- `http_transport.New()` — WASM implementation, registers handlers as `window` globals (`"api/echo"` → `window.api_echo`)
- `docs/http_transport.js` — JS client, checks `window[name]` first, falls back to `fetch`
- Zero external dependencies
- Go 1.18+ (generics), build tags: `!js` for server, `js && wasm` for browser

## DISCLAIMER

This project was implemented with assistance from Claude Sonnet 4.6 (claude-sonnet-4-6) by Anthropic.
