package prepare

import (
	"net/http"

	"github.com/jfyne/live"
	"github.com/tinygodsdev/tinycooksweb/internal/config"
)

func Store(cfg *config.Config) *live.CookieStore {
	store := live.NewCookieStore(cfg.LiveSessionName, []byte(cfg.Secret))
	store.Store.Options.SameSite = http.SameSiteLaxMode
	store.Store.MaxAge(cfg.Expire)
	store.Store.Options.Path = "/"
	store.Store.Options.HttpOnly = true
	store.Store.Options.Secure = !(cfg.Env == "dev")
	return store
}
