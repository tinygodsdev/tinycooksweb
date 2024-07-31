package app

import (
	"github.com/tinygodsdev/datasdk/pkg/logger"
	"github.com/tinygodsdev/tinycooksweb/internal/config"
	"github.com/tinygodsdev/tinycooksweb/pkg/locale"
)

type (
	App struct {
		Cfg     *config.Config
		log     logger.Logger
		locales []string // locale count is not very big so no need to have map
		Errors  []error
	}
)

func New(
	cfg *config.Config,
	logger logger.Logger,
) (*App, error) {
	app := App{
		Cfg:     cfg,
		log:     logger,
		locales: locale.List(),
	}

	return &app, nil
}

func (a *App) IsDev() bool {
	return a.Cfg.Dev()
}

func (a *App) AddError(err error) {
	a.Errors = append(a.Errors, err)
}
