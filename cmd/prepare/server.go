package prepare

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jfyne/live"

	"github.com/tinygodsdev/tinycooksweb/internal/config"
	"github.com/tinygodsdev/tinycooksweb/internal/handler"
)

func Mux(cfg *config.Config, store live.HttpSessionStore, h *handler.Handler) *mux.Router {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(h.NotFoundRedirect)
	// main handler
	sr := r.PathPrefix("/").Subrouter()
	sr.Handle(fmt.Sprintf("/recipe/{%s}", handler.ParamRecipeSlug), live.NewHttpHandler(store, h.Recipe()))
	sr.Handle(fmt.Sprintf("/tag/{%s}", handler.ParamTagSlug), live.NewHttpHandler(store, h.Tag()))
	sr.Handle(fmt.Sprintf("/ingredient/{%s}", handler.ParamIngredientSlug), live.NewHttpHandler(store, h.Ingredient()))
	sr.Handle(fmt.Sprintf("/equipment/{%s}", handler.ParamEquipmentSlug), live.NewHttpHandler(store, h.Equipment()))
	sr.Handle("/about", live.NewHttpHandler(store, h.About()))
	// sr.Handle("/profile", live.NewHttpHandler(store, h.Profile()))
	sr.Handle("/privacy", live.NewHttpHandler(store, h.Privacy()))
	sr.Handle("/terms", live.NewHttpHandler(store, h.Terms()))
	sr.Handle("/catalog", live.NewHttpHandler(store, h.Catalog()))
	// sr.Handle("/admin", live.NewHttpHandler(store, h.Admin()))
	// sr.Handle("/status", live.NewHttpHandler(store, h.Status()))
	sr.Handle("/404", live.NewHttpHandler(store, h.NotFound()))
	sr.Handle("/", live.NewHttpHandler(store, h.Home()))

	// index
	r.HandleFunc("/robots.txt", h.RobotsHandler())
	r.HandleFunc("/sitemap.xml", h.SitemapHandler())

	// live scripts
	r.Handle("/live.js", live.Javascript{})
	r.Handle("/auto.js.map", live.JavascriptMap{})

	// static
	r.HandleFunc("/favicon.ico", faviconHandler)
	r.HandleFunc("/android-chrome-192x192.png", androidChrome192x192Handler)
	r.HandleFunc("/android-chrome-512x512.png", androidChrome512x512Handler)
	r.HandleFunc("/apple-touch-icon.png", appleTouchIconHandler)
	r.HandleFunc("/favicon-16x16.png", favicon16x16Handler)
	r.HandleFunc("/favicon-32x32.png", favicon32x32Handler)

	r.HandleFunc("/static/css/styles.css", stylesHandler)

	return r
}

func Server(cfg *config.Config, handler *mux.Router) *http.Server {
	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      handler,
	}

	return server
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/media/favicon.ico")
}

func stylesHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/dist/css/styles.css")
}

func androidChrome192x192Handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/media/android-chrome-192x192.png")
}

func androidChrome512x512Handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/media/android-chrome-512x512.png")
}

func appleTouchIconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/media/apple-touch-icon.png")
}

func favicon16x16Handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/media/favicon-16x16.png")
}

func favicon32x32Handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/media/favicon-32x32.png")
}
