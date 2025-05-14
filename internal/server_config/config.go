package server_config

import (
	"flag"
	"os"
	"time"
)

type Config struct {
	RunAddr          string
	LogLevel         string
	DBDSN            string
	SecretKey        string
	TokenExp         time.Duration
	SecretKeyForData string
}

func NewConfig() (*Config, error) {
	config := &Config{
		TokenExp: time.Hour * 1344, //2 месяца не истекает авторизация
	}
	config.SetValues()
	return config, nil
}

func (c *Config) SetValues() {
	// регистрируем переменную flagRunAddr как аргумент -a со значением по умолчанию localhost:8080
	flag.StringVar(&c.RunAddr, "a", "localhost:8080", "Set listening address and port for server")
	// регистрируем уровень логирования
	flag.StringVar(&c.LogLevel, "l", "debug", "logger level")
	// принимаем строку подключения к базе данных
	//flag.StringVar(&c.DBDSN, "d", "postgresql://gophkeeper:gophkeeper@localhost:5432/gophkeeper?sslmode=disable", "postgres database")
	//flag.StringVar(&c.DBDSN, "d", "postgresql://user1:Newpass34,@c-c9q5978bins6fbk66o32.rw.mdb.yandexcloud.net:6432/db1", "postgres database")
	flag.StringVar(&c.DBDSN, "d", "postgresql://postgres:j0Wam3ibcT4KnGWUWuabEpuUmzL@212.193.48.196:5432/template1", "postgres database")
	// принимаем секретный ключ сервера для авторизации
	flag.StringVar(&c.SecretKey, "s", "e4853f5c4810101e88f1898db21c15d3", "server's secret key for authorization")
	// принимаем секретный ключ сервера для расшифровки
	flag.StringVar(&c.SecretKeyForData, "sd", "TuUdlQmYyD1DTaiGVV31ipyWnbKa0jUD", "secret key for data encrypting")

	if envRunAddr := os.Getenv("RUN_ADDRESS"); envRunAddr != "" {
		c.RunAddr = envRunAddr
	}
	if envLogLevel := os.Getenv("LOG_LEVEL"); envLogLevel != "" {
		c.LogLevel = envLogLevel
	}
	if envDBDSN := os.Getenv("DATABASE_URI"); envDBDSN != "" {
		c.DBDSN = envDBDSN
	}
	if envSecretKey := os.Getenv("SECRET_KEY"); envSecretKey != "" {
		c.SecretKey = envSecretKey
	}
	if envSecretKeyForData := os.Getenv("SECRET_KEY_FOR_DATA"); envSecretKeyForData != "" {
		c.SecretKeyForData = envSecretKeyForData
	}
}
