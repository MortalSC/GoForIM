package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"gopkg.in/yaml.v3"
)

var DB *sql.DB

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
}

func LoadConfig() (*Config, error) {
	file, err := os.ReadFile("etc/config.yaml")
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func InitDB() {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config: %v", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database: %v", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("Failed to ping database: %v", err)
	}

	log.Println("Successfully connected to database")
}
