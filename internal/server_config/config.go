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
	flag.StringVar(&c.DBDSN, "d", "postgresql://gophkeeper:gophkeeper@localhost:5432/gophkeeper?sslmode=disable", "postgres database")
	// принимаем секретный ключ сервера для авторизации
	flag.StringVar(&c.SecretKey, "s", "", "server's secret key for authorization")
	// принимаем секретный ключ сервера для расшифровки
	flag.StringVar(&c.SecretKeyForData, "sd", "", "secret key for data encrypting")

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
