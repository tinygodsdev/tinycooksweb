package timer

import (
	"time"

	"github.com/tinygodsdev/datasdk/pkg/logger"
)

func Info(name string, logger logger.Logger, args ...interface{}) func() {
	start := time.Now()
	return func() {
		args = append(args, "took", time.Since(start).String())
		logger.Info(name, args...)
	}
}
