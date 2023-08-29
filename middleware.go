package dataprovider

import (
	"log"
	"net/http"
	"strings"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"github.com/devetek/caddy-dataprovider/controllers"
	"github.com/devetek/caddy-dataprovider/utils"
)

const moduleName = "data_provider"

func init() {
	caddy.RegisterModule(Middleware{})
	httpcaddyfile.RegisterHandlerDirective(moduleName, parseCaddyfile)
	// httpcaddyfile.RegisterGlobalOption(moduleName, parseCaddyfileGlobalOption)
}

type Module struct {
	ModuleName string   `json:"module_name,omitempty"`
	Schema     []string `json:"schema,omitempty"`
}

type Middleware struct {
	Module Module `json:"module,omitempty"`
}

// CaddyModule returns the Caddy module information.
func (Middleware) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.data_provider",
		New: func() caddy.Module { return new(Middleware) },
	}
}

// Provision implements caddy.Provisioner.
// https://pkg.go.dev/github.com/caddyserver/caddy/v2#Provisioner
// Invoke registered modules in the Caddy configuration
func (m *Middleware) Provision(ctx caddy.Context) error {
	return nil
}

// Validate implements caddy.Validator.
// Call after provisioning to make sure the module
func (m *Middleware) Validate() error {
	return nil
}

// ServeHTTP implements caddyhttp.MiddlewareHandler.
func (m Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {

	data := controllers.GetDataBlogHome()

	repl := r.Context().Value(caddy.ReplacerCtxKey).(*caddy.Replacer)

	r.Header.Add("Kratos-Edge-Data", utils.RemoveNewLine(data))

	// this will be used to determine dynamic params from the URL
	// used to query to the origin API / graphql
	log.Printf("==================== || ====================")
	matchRegex, _ := repl.GetString("http.regexp.0")
	dynamicStr := strings.Replace(r.URL.Path, matchRegex, "", -1)

	// log.Println(data)
	log.Println(dynamicStr)
	log.Printf("==================== || ====================")

	// log.Printf("===================")
	// log.Println(repl)
	// log.Println(repl.Get("http.regexp.0"))
	// log.Println(repl.Get("http.regexp.0.name"))
	// log.Println(repl.GetString("http.regexp.0"))
	// log.Printf("===================")

	return next.ServeHTTP(w, r)
}

// UnmarshalCaddyfile implements caddyfile.Unmarshaler.
func (m *Middleware) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	return nil
}

// func parseCaddyfileGlobalOption(h *caddyfile.Dispenser, _ interface{}) (interface{}, error) {
// 	souinApp := new(SouinApp)
// 	cfg := &Configuration{
// 		DefaultCache: &DefaultCache{
// 			AllowedHTTPVerbs: make([]string, 0),
// 			Distributed:      false,
// 			Headers:          []string{},
// 			TTL: configurationtypes.Duration{
// 				Duration: 120 * time.Second,
// 			},
// 			DefaultCacheControl: "",
// 			CacheName:           "",
// 		},
// 		URLs: make(map[string]configurationtypes.URL),
// 	}

// 	err := parseConfiguration(cfg, h, true)

// 	souinApp.DefaultCache = cfg.DefaultCache
// 	souinApp.API = cfg.API
// 	souinApp.CacheKeys = cfg.CacheKeys
// 	souinApp.LogLevel = cfg.LogLevel

// 	return httpcaddyfile.App{
// 		Name:  moduleName,
// 		Value: caddyconfig.JSON(souinApp, nil),
// 	}, err
// }

// parseCaddyfile unmarshals tokens from h into a new Middleware.
func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	var m Middleware
	err := m.UnmarshalCaddyfile(h.Dispenser)
	return m, err
}

// Interface guards
var (
	_ caddy.Provisioner           = (*Middleware)(nil)
	_ caddy.Validator             = (*Middleware)(nil)
	_ caddyhttp.MiddlewareHandler = (*Middleware)(nil)
	_ caddyfile.Unmarshaler       = (*Middleware)(nil)
)
