package config

import (
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port          string
	DBPostgresURL string
	JWTSecretKey  string
}

const configFile = "./internal/config/envs/dev.env"

func NewConfig() (*Config, error) {
	cfg := Config{}

	cfg.setDefualtValues()
	cfg.loadEnv()
	cfg.getFlagConfig()

	return &cfg, nil
}

func (c *Config) setDefualtValues() {
	c.Port = ":50001"
	c.DBPostgresURL = "postgres://<USER>:<PASSWORD>@<HOST>:<PORT>/auth_db"
	c.JWTSecretKey = "JWT SECRET KEY"
}

func (c *Config) loadEnv() {
	err := godotenv.Load(configFile)
	if err != nil {
		log.Println(err)
	}

	port, ok := os.LookupEnv("PORT")
	if ok {
		c.Port = port
	}

	dbURL, ok := os.LookupEnv("DB_URL")
	if ok {
		c.DBPostgresURL = dbURL
	}

	jwtKey, ok := os.LookupEnv("JWT_SECRET_KEY")
	if ok {
		c.JWTSecretKey = jwtKey
	}
}

func (c *Config) getFlagConfig() {
	flag.StringVar(&c.Port, "p", c.Port, "API Gateway service port, example :30001")
	flag.StringVar(&c.DBPostgresURL, "db", c.DBPostgresURL, "DB Postgres URL, example postgres://<USER>:<PASSWORD>@<HOST>:<PORT>/auth_db")
	flag.StringVar(&c.JWTSecretKey, "j", c.JWTSecretKey, "JWT secret key, example secretKey")
	flag.Parse()
}
