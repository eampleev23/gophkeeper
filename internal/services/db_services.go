package services

import (
	"github.com/eampleev23/gophkeeper/internal/mlg"
	"github.com/eampleev23/gophkeeper/internal/myauth"
	"github.com/eampleev23/gophkeeper/internal/server_config"
	"github.com/eampleev23/gophkeeper/internal/store"
)

type DBServices struct {
	s  store.Store
	c  *server_config.Config
	l  *mlg.ZapLog
	au myauth.Authorizer
}

func NewDBServices(s store.Store, c *server_config.Config, l *mlg.ZapLog, au myauth.Authorizer) Services {
	dbServices := &DBServices{
		s:  s,
		c:  c,
		l:  l,
		au: au,
	}
	return dbServices
}
