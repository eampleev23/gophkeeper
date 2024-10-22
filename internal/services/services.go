package services

import (
	"github.com/eampleev23/gophkeeper/internal/mlg"
	"github.com/eampleev23/gophkeeper/internal/myauth"
	"github.com/eampleev23/gophkeeper/internal/server_app"
	"github.com/eampleev23/gophkeeper/internal/store"
)

type Services struct {
	s  store.Store
	c  *server_app.Config
	l  *mlg.ZapLog
	au myauth.Authorizer
}

func NewServices(s store.Store, c *server_app.Config, l *mlg.ZapLog, au myauth.Authorizer) *Services {
	services := &Services{
		s:  s,
		c:  c,
		l:  l,
		au: au,
	}
	return services
}
