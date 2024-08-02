package app

import (
	"github.com/tinygodsdev/tinycooksweb/internal/util/timer"
)

func (a *App) Timer(method string, args ...interface{}) func() {
	name := "App." + method
	return timer.Info(name, a.log, args...)
}
