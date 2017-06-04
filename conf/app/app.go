package app

import (
	mcache "github.com/go-macaron/cache"
	"github.com/go-macaron/gzip"
	"github.com/go-macaron/i18n"
	"github.com/go-macaron/jade"
	"github.com/go-macaron/session"
	"github.com/go-macaron/toolbox"
	"github.com/novatrixtech/aguiaapi/conf"
	"github.com/novatrixtech/aguiaapi/handler"
	"github.com/novatrixtech/aguiaapi/lib/cache"
	"github.com/novatrixtech/aguiaapi/lib/context"
	"github.com/novatrixtech/aguiaapi/lib/template"
	"gopkg.in/macaron.v1"
)

// SetupMiddlewares
func SetupMiddlewares(app *macaron.Macaron) {
	app.Use(macaron.Logger())
	app.Use(macaron.Recovery())
	app.Use(gzip.Gziper())
	app.Use(toolbox.Toolboxer(app, toolbox.Options{
		HealthCheckers: []toolbox.HealthChecker{
			new(handler.AppChecker),
		},
	}))
	app.Use(macaron.Static("public"))
	app.Use(i18n.I18n(i18n.Options{
		Directory: "locale",
		Langs:     []string{"pt-BR", "en-US"},
		Names:     []string{"Português do Brasil", "American English"},
	}))
	app.Use(jade.Renderer(jade.Options{
		Directory: "public/templates",
		Funcs:     template.FuncMaps(),
	}))
	app.Use(macaron.Renderer(macaron.RenderOptions{
		Directory: "public/templates",
		Funcs:     template.FuncMaps(),
	}))
	app.Use(mcache.Cacher(
		cache.Option(conf.Cfg.Section("").Key("cache_adapter").Value()),
	))
	app.Use(session.Sessioner())
	app.Use(context.Contexter())
}

// SetupRoutes
func SetupRoutes(app *macaron.Macaron) {
	app.Get("", func() string {
		return "Mercurius Works!"
	})

	app.Group("/api/v1", func() {
		app.Get("/categorias", handler.GetCategoria)
		app.Get("/contratos", handler.GetContrato)
	})
}
