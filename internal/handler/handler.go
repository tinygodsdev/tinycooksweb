package handler

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"time"

	"github.com/bradfitz/iter"
	"github.com/jfyne/live"

	"github.com/tinygodsdev/datasdk/pkg/logger"
	"github.com/tinygodsdev/tinycooksweb/internal/app"
	"github.com/tinygodsdev/tinycooksweb/pkg/locale"
)

const (
	// views
	view404     = "404"
	viewAbout   = "about"
	viewAdmin   = "admin"
	viewTest    = "test"
	viewResult  = "result"
	viewHome    = "home"
	viewProfile = "profile"
	viewPrivacy = "privacy"
	viewTerms   = "terms"
	viewStatus  = "status"
	// events (common)
	eventCloseError   = "close-error-notification"
	eventCloseMessage = "close-message-notification"
	eventToggleDark   = "toggle-dark"
	// params (common)
	paramLocale = "locale"
	// context
	userCtxKeyValue   = "user"
	localeCtxKeyValue = "locale"

	DefaultDisplayTime = time.RFC822
)

type (
	Handler struct {
		app *app.App
		log logger.Logger
		t   string // template path
		ui  map[string]*locale.UITranslation
	}

	CommonInstance struct {
		Env               string
		Domain            string
		Session           string
		Error             error
		Message           *string
		CurrentView       string
		Version           string
		GoogleAnalyticsID string
		UI                *locale.UITranslation
		ui                map[string]*locale.UITranslation
		locale            string
	}

	Constants struct {
		// to have constants in templates
		IntroStatus     string
		QuestionsStatus string
		FinishStatus    string
		ResultStatus    string
		MethodSten      string
		MethodPerc      string
		MethodMean      string
		MethodSum       string
	}

	contextKey struct {
		name string
	}
)

var userCtxKey = &contextKey{userCtxKeyValue}
var localeCtxKey = &contextKey{localeCtxKeyValue}
var funcMap = template.FuncMap{
	"N":     iter.N,
	"Plus1": func(i int) int { return i + 1 },
	"Sum": func(data ...float64) float64 {
		var res float64
		for _, n := range data {
			res += n
		}
		return res
	},
	"Sub": func(f1, f2 float64) float64 {
		return f1 - f2
	},
	"Mean": func(data ...float64) float64 {
		if len(data) == 0 {
			return 0
		}
		var sum float64
		for _, n := range data {
			sum += n
		}
		return sum / float64(len(data))
	},
	"Perc": func(min, max, v float64) float64 {
		if max == min {
			return 0
		}
		return (v - min) / (max - min)
	},
	"DerefInt": func(i *int) int {
		if i == nil {
			return 0
		}
		return *i
	},
	"DisplayTime": func(t time.Time) string {
		return t.Format(DefaultDisplayTime)
	},
	"DisplayTechTime": func(t time.Time) string {
		return t.Format("2006-01-02 15:04:05.000 MST")
	},
	"Since": func(t time.Time) time.Duration {
		return time.Since(t)
	},
	"UILocales": func() []string {
		return locale.List()
	},
	"LocaleParam": func(loc string) string {
		if loc == locale.Default() {
			return "/"
		}
		return fmt.Sprintf("?locale=%s", loc)
	},
	"FormatDuration": func(d time.Duration) string {
		z := time.Unix(0, 0).UTC()
		return z.Add(d).Format("4:05")
	},
}

func NewHandler(
	app *app.App,
	logger logger.Logger,
	t string,
) *Handler {
	return &Handler{
		app: app,
		log: logger,
		t:   t,
		ui:  initTranslationMap(),
	}
}

func initTranslationMap() map[string]*locale.UITranslation {
	m := make(map[string]*locale.UITranslation)
	for _, l := range locale.List() {
		m[l] = locale.NewUITranslation(l)
	}

	return m
}

func (h *Handler) NewConstants() *Constants {
	return &Constants{}
}

func (h *Handler) NewCommon(s live.Socket, currentView string) *CommonInstance {
	c := &CommonInstance{
		Env:               h.app.Cfg.Env,
		Domain:            h.app.Cfg.HTTPHost,
		Session:           fmt.Sprint(s.Session()),
		Error:             nil,
		Message:           nil,
		CurrentView:       currentView,
		Version:           h.app.Cfg.Version,
		GoogleAnalyticsID: h.app.Cfg.GoogleAnalyticsID,
		ui:                h.ui,
		locale:            locale.Default(), // it's private because changing requires additional logic
	}

	c.SetLocale(c.locale)
	return c
}

func (h *Handler) url404() *url.URL {
	u, _ := url.Parse("/404")
	return u
}

func (h *Handler) HandleError(ctx context.Context, err error) {
	h.log.Error("got bad request", "err", err)
	w := live.Writer(ctx)
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("bad request: " + err.Error()))
}

func (h *Handler) NotFoundRedirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, h.url404().String(), http.StatusTemporaryRedirect)
}

func (c *CommonInstance) getTranslation(loc string) *locale.UITranslation {
	trans, ok := c.ui[loc]
	if !ok {
		return c.ui[locale.Default()]
	}

	return trans
}

func (c *CommonInstance) CloseError() {
	c.Error = nil
}

func (c *CommonInstance) CloseMessage() {
	c.Message = nil
}

func (c *CommonInstance) SetLocale(l string) {
	if !locale.IsValid(l) {
		l = locale.Default()
	}
	c.locale = l
	c.UI = c.getTranslation(l)
}

func (c *CommonInstance) Locale() string {
	return c.locale
}

func localeFromCtx(ctx context.Context) string {
	loc, ok := ctx.Value(localeCtxKey).(string)
	if !ok {
		return locale.Default()
	}
	return loc
}

func (c *CommonInstance) fromContext(ctx context.Context) {
	c.locale = localeFromCtx(ctx)
}
