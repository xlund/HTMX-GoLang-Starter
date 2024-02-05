package main

import (
	"context"
	"fmt"
	"go-starter/routes"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	ory "github.com/ory/client-go"
)

type App struct {
	router *chi.Mux
	ory    *ory.APIClient
}

type AppSettings struct {
	Proxy string
}

func NewApp(settings AppSettings) App {
	r := createRouter()
	oc := createOryClient(settings.Proxy)
	app := App{
		router: r,
		ory:    oc,
	}
	return app
}

func (app *App) Start(addr string) error {
	app.ServeRoutes()
	err := http.ListenAndServe(addr, app.router)
	return err
}

func (app *App) ServeRoutes() {
	app.router.Get("/", routes.Home)
	app.router.Get("/user", app.sessionMiddleware(routes.UserHandler))
}

func createOryClient(proxy string) *ory.APIClient {
	c := ory.NewConfiguration()
	c.Servers = ory.ServerConfigurations{
		{
			URL: proxy,
		},
	}
	return ory.NewAPIClient(c)
}

func createRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	handleStatic(r)
	return r
}

func handleStatic(r *chi.Mux) {
	prefix := "/static/"
	dir := http.Dir("html/static")
	fs := http.FileServer(dir)

	r.Handle("/static/*", http.StripPrefix(prefix, fs))
}

func (app *App) sessionMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// do something with the request
		log.Default().Println("sessionMiddleware: before")

		cookies := r.Header.Get("Cookie")

		session, _, err := app.ory.FrontendAPI.ToSession(r.Context()).Cookie(cookies).Execute()
		if (err != nil && session == nil) || (err == nil && !*session.Active) {
			// this will redirect the user to the managed Ory Login UI
			redirectUrl := fmt.Sprintf("%v/.ory/self-service/login/browser", *proxy)
			println(redirectUrl)
			log.Default().Println("sessionMiddleware: redirecting to", redirectUrl)
			http.Redirect(w, r, redirectUrl, http.StatusSeeOther)
			return
		}
		ctx := withCookies(r.Context(), cookies)
		ctx = withSession(ctx, session)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func withCookies(ctx context.Context, cookies string) context.Context {
	return context.WithValue(ctx, "req.cookies", cookies)
}

func withSession(ctx context.Context, session *ory.Session) context.Context {
	return context.WithValue(ctx, "req.session", session)
}
