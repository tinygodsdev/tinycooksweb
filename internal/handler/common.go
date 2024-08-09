package handler

import (
	"context"
	"fmt"

	"github.com/jfyne/live"
	"github.com/tinygodsdev/tinycooksweb/pkg/locale"
)

type (
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
)

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
