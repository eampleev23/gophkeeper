package main

import (
	"fmt"
	"github.com/eampleev23/gophkeeper/internal/cnf"
	"github.com/eampleev23/gophkeeper/internal/handlers"
	"github.com/eampleev23/gophkeeper/internal/mlg"
	"github.com/eampleev23/gophkeeper/internal/myauth"
	"github.com/eampleev23/gophkeeper/internal/services"
	"github.com/eampleev23/gophkeeper/internal/store"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"log"
	"net/http"
)

/*Сурвер GophKeeper позволяет клиентам пользователя надёжно
и безопасно хранить логины, пароли, бинарные данные и прочую приватную информацию.
А также отвечает за синхронизацию между разными клиентами пользователя

Реализует следующую бизнес-логику:
регистрация, аутентификация и авторизация пользователей;
хранение приватных данных;
синхронизация данных между несколькими авторизованными клиентами одного владельца;
передача приватных данных владельцу по запросу.
*/

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	c, err := cnf.NewConfig()
	if err != nil {
		return fmt.Errorf("failed to initialize a new config: %w", err)
	}
	mL, err := mlg.NewZapLogger(c.LogLevel)
	if err != nil {
		return fmt.Errorf("failed to initialize a new logger: %w", err)
	}
	mL.ZL.Debug("Logger success created..")
	au, err := myauth.Initialize(c, mL)
	if err != nil {
		return fmt.Errorf("failed to initialize a new authorizer: %w", err)
	}

	s, err := store.NewStorage(c, mL)
	if err != nil {
		return fmt.Errorf("failed to initialize a new store: %w", err)
	}
	if len(c.DBDSN) != 0 {
		// Отложенно закрываем соединение с бд.
		defer func() {
			if err := s.DBConnClose(); err != nil {
				mL.ZL.Info("store failed to properly close the DB connection")
			}
		}()
	}
	serv := services.NewServices(s, c, mL, *au)
	h, err := handlers.NewHandlers(s, c, mL, *au, *serv)
	if err != nil {
		return fmt.Errorf("handlers constructor's error: %w", err)
	}
	mL.ZL.Info("Running server", zap.String("address", c.RunAddr))

	r := chi.NewRouter()
	r.Use(mL.RequestLogger)
	r.Post("/api/user/register", h.Register)
	r.Post("/api/user/login", h.Authentication)
	r.Post("/api/user/add-login-password", h.AddLoginPasswordHandler)
	r.Get("/api/user/get-data-items", h.GetDataItems)
	err = http.ListenAndServe(c.RunAddr, r)
	if err != nil {
		return fmt.Errorf("ошибка ListenAndServe: %w", err)
	}
	return nil
}
