package handlers

import (
	"github.com/eampleev23/gophkeeper/internal/mlg"
	"github.com/eampleev23/gophkeeper/internal/myauth"
	"github.com/eampleev23/gophkeeper/internal/server_config"
	"github.com/eampleev23/gophkeeper/internal/services"
	"github.com/eampleev23/gophkeeper/internal/store"
	"net/http"
)

type Handlers struct {
	s    store.Store
	c    *server_config.Config
	l    *mlg.ZapLog
	au   myauth.Authorizer
	serv services.Services
}

func NewHandlers(
	s store.Store,
	c *server_config.Config,
	l *mlg.ZapLog,
	au myauth.Authorizer,
	serv services.Services) (
	*Handlers,
	error) {
	return &Handlers{
		s:    s,
		c:    c,
		l:    l,
		au:   au,
		serv: serv,
	}, nil
}

func (h *Handlers) GetUserID(r *http.Request) (userID int, isAuth bool, err error) {
	h.l.ZL.Debug("GetUserID started.. ")
	cookie, err := r.Cookie("token")
	if err != nil {
		return 0, false, nil //nolint:nilerr // нужно будет исправить логику
	}
	userID, err = h.au.GetUserID(cookie.Value)
	if err != nil {
		return 0, false, nil //nolint:nilerr // нужно будет исправить логику
	}
	return userID, true, nil
}
