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
		Email    Email
		Storage  Storage
		Gemini   Gemini
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

	Email struct {
		Host     string
		Sender   string
		Port     int
		Email    string
		Password string
	}

	Storage struct {
		ApiKey    string
		ProjectID string
		Bucket    string
	}

	Gemini struct {
		ApiKey string
		Model  string
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

	emailPort, err := strconv.Atoi(os.Getenv("EMAIL_SMTP_PORT"))
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

	storage := Storage{
		ApiKey:    os.Getenv("STORAGE_API_KEY"),
		ProjectID: os.Getenv("STORAGE_PROJECT_ID"),
		Bucket:    os.Getenv("STORAGE_BUCKET_ID"),
	}

	email := Email{
		Host:     os.Getenv("EMAIL_SMTP_HOST"),
		Port:     emailPort,
		Email:    os.Getenv("EMAIL_SMTP_EMAIL"),
		Password: os.Getenv("EMAIL_SMTP_PASSWORD"),
		Sender:   os.Getenv("EMAIL_SMTP_SENDER"),
	}

	gemini := Gemini{
		ApiKey: os.Getenv("GEMINI_API_KEY"),
		Model:  os.Getenv("GEMINI_MODEL"),
	}

	env := &Env{
		App:      app,
		Database: database,
		Cache:    cache,
		Email:    email,
		Storage:  storage,
		Gemini:   gemini,
	}

	return env, nil
}
