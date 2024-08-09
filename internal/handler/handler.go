package handler

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/jfyne/live"

	"github.com/tinygodsdev/datasdk/pkg/logger"
	"github.com/tinygodsdev/tinycooksweb/internal/app"
	"github.com/tinygodsdev/tinycooksweb/pkg/locale"
)

const (
	// views
	view404     = "404"
	viewAbout   = "about"
	viewRecipe  = "recipe"
	viewHome    = "home"
	viewProfile = "profile"
	viewPrivacy = "privacy"
	viewTerms   = "terms"
	viewStatus  = "status"
	// events (common)
	eventCloseError   = "close-error-notification"
	eventCloseMessage = "close-message-notification"
	// params (common)
	paramLocale = "locale"
	// context
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
	}

	contextKey struct {
		name string
	}
)

var localeCtxKey = &contextKey{localeCtxKeyValue}

func NewHandler(
	app *app.App,
	logger logger.Logger,
	t string,
) (*Handler, error) {
	tm, err := initTranslationMap()
	if err != nil {
		return nil, err
	}

	return &Handler{
		app: app,
		log: logger,
		t:   t,
		ui:  tm,
	}, nil
}

func initTranslationMap() (map[string]*locale.UITranslation, error) {
	m := make(map[string]*locale.UITranslation)
	for _, l := range locale.List() {
		trans := locale.NewUITranslation(l)
		if err := trans.Validate(); err != nil {
			return nil, err
		}

		m[l] = trans
	}

	return m, nil
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
