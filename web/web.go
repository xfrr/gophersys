package web

import (
	"embed"
	"html/template"
	"io/fs"

	"mime"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/xfrr/gophersys/pkg/bus"
	"github.com/xfrr/gophersys/pkg/logger"
)

// embed static files
var (
	//go:embed static/*
	staticFS embed.FS
)

// load embedded static files
var (
	fsys            = fs.FS(staticFS)
	staticAssets, _ = fs.Sub(fsys, "static")
)

func init() {
	_ = mime.AddExtensionType(".js", "text/javascript")
}

func NewApp(cmdbus bus.Bus, queryBus bus.Bus, logger logger.Logger) (*App, error) {
	app := &App{
		cmdbus:   cmdbus,
		queryBus: queryBus,
		logger:   logger,
	}

	fm := template.FuncMap{
		"getGophers":    app.getGophers,
		"deleteGophers": app.deleteGophers,
		"createGopher":  app.createGopher,
		"updateGopher":  app.updateGopher,
	}

	if err := app.loadTemplates(fm); err != nil {
		return nil, err
	}

	return app, nil
}

// App represents the web application instance
type App struct {
	cmdbus   bus.Bus
	queryBus bus.Bus
	logger   logger.Logger
}

func (a App) ListenAndServe(port string, certFile, keyFile string) error {
	if certFile != "" && keyFile != "" {
		return http.ListenAndServeTLS(":"+port, certFile, keyFile, a.newRouter())
	}

	return http.ListenAndServe(":"+port, a.newRouter())
}

func (a App) newRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.Heartbeat("/health"))

	// serve index page
	r.Get("/", indexHandler())

	// serve static files
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.FS(staticAssets))))
	return r
}
