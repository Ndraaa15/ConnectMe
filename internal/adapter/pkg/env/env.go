package env

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type (
	Env struct {
		App      App
		Database Database
		Cache    Cache
	}

	App struct {
		Name    string
		Port    string
		Address string
	}

	Database struct {
		Address  string
		Port     string
		User     string
		Password string
		Name     string
		SSLMode  string
	}

	Cache struct {
		Address  string
		Port     string
		Password string
		DB       int
		Username string
	}
)

func NewEnv() (*Env, error) {
	if err := godotenv.Load(); err != nil {
		return &Env{}, err
	}

	cacheDB, err := strconv.Atoi(os.Getenv("CACHE_DB"))
	if err != nil {
		return &Env{}, err
	}

	app := App{
		Name:    os.Getenv("APP_NAME"),
		Port:    os.Getenv("APP_PORT"),
		Address: os.Getenv("APP_ADDRESS"),
	}

	database := Database{
		Address:  os.Getenv("DATABASE_ADDRESS"),
		Port:     os.Getenv("DATABASE_PORT"),
		User:     os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		Name:     os.Getenv("DATABASE_NAME"),
		SSLMode:  os.Getenv("DATABASE_SSL_MODE"),
	}

	cache := Cache{
		Address:  os.Getenv("CACHE_ADDRESS"),
		Port:     os.Getenv("CACHE_PORT"),
		Password: os.Getenv("CACHE_PASSWORD"),
		DB:       cacheDB,
		Username: os.Getenv("CACHE_USERNAME"),
	}

	env := &Env{
		App:      app,
		Database: database,
		Cache:    cache,
	}

	return env, nil
}
