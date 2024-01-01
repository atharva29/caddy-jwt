package jwt

import (
	"fmt"
	"net/http"

	// _ "github.com/example/jwt"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	caddy.RegisterModule(JWTAuth{})
	httpcaddyfile.RegisterDirective("jwtauth", nil)
}

type JWTAuth struct{}

func (JWTAuth) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.jwtauth",
		New: func() caddy.Module { return new(JWTAuth) },
	}
}

func (h *JWTAuth) Provision(ctx caddy.Context) error {
	return nil
}

func (h *JWTAuth) Validate() error {
	return nil
}

func (h JWTAuth) ServeHTTP(w http.ResponseWriter, r *http.Request, f caddyhttp.Handler) error {
	fmt.Println("**********  Setting Header ****************************")
	// Add the custom header
	w.Header().Set("message", "Hello world")

	// Reverse proxy to localhost:5000
	// proxy := reverseproxy.New(&http.Server{
	// 	Addr: "localhost:5000",
	// })
	return f.ServeHTTP(w, r)
}

// Interface guards
var (
	_ caddyhttp.MiddlewareHandler = (*JWTAuth)(nil)
)
