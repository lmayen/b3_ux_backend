package config

import (
	"b3_ux_backend/internal/fsutils"
	"fmt"
	"path/filepath"
	"time"

	"github.com/olekukonko/errors"
)

var CONFIG *Config

type AppConfig struct {
	Name       string `json:"name"`
	Version    string `json:"version"`
	RootDir    string `json:"root_dir"`
	LogDir     string `json:"log_dir"`
	ImageDir   string `json:"image_dir"`
	SqlitePath string `json:"sqlite_path"`
}

type HTTPConfig struct {
	Host             string `json:"host"`
	Port             int    `json:"port"`
	WebSocketEnabled bool   `json:"webSocketEnabled"`
	SecureCookies    bool   `json:"secureCookies"`
}

type CORSConfig struct {
	Enabled          bool     `json:"enabled"`
	AllowedOrigins   []string `json:"allowed_origins"`
	AllowedMethods   []string `json:"allowed_methods"`
	AllowedHeaders   []string `json:"allowed_headers"`
	AllowCredentials bool     `json:"allow_credentials"`
}

type DatabaseConfig struct {
	MaxOpenConns    int           `json:"max_open_conns"`
	MaxIdleConns    int           `json:"max_idle_conns"`
	ConnMaxLifetime time.Duration `json:"conn_max_lifetime"` // e.g. 1h
	PingTimeout     time.Duration `json:"ping_timeout"`      // e.g. 3s
	Name            string        `json:"name"`
}

type Config struct {
	App      AppConfig      `json:"app"`
	HTTP     HTTPConfig     `json:"http"`
	CORS     CORSConfig     `json:"cors"`
	Database DatabaseConfig `json:"database"`
}

func (c *Config) Validate() error {
	if c.App.Name == "" {
		return errors.New("app.name must not be empty")
	}

	if c.HTTP.Port <= 0 || c.HTTP.Port > 65535 {
		return fmt.Errorf("invalid http.port: %d", c.HTTP.Port)
	}

	if c.Database.ConnMaxLifetime <= 0 {
		return errors.New("database.conn_max_lifetime must be > 0")
	}

	if c.Database.PingTimeout <= 0 {
		return errors.New("database.ping_timeout must be > 0")
	}

	if c.Database.MaxIdleConns <= 0 {
		return errors.New("database.max_idle_conns must be > 0")
	}

	if c.Database.MaxOpenConns <= 0 {
		return errors.New("database.max_open_conns must be > 0")
	}

	if c.Database.ConnMaxLifetime <= 0 {
		return errors.New("database.conn_max_lifetime must be > 0")
	}

	if c.Database.Name == "" {
		return errors.New("database.name must not be empty")
	}

	return nil
}

func InitConfig() (*Config, error) {
	appName := "b3_ux_backend"
	root, err := fsutils.AppDir(appName)
	if err != nil {
		return nil, errors.Wrapf(err, "InitConfig")
	}

	err = fsutils.EnsureDir(root)
	if err != nil {
		return nil, errors.Wrapf(err, "InitConfig")
	}

	logDirectory := filepath.Join(root, "log")
	err = fsutils.EnsureDir(logDirectory)
	if err != nil {
		return nil, errors.Wrapf(err, "InitConfig")
	}

	imgDirectory := filepath.Join(root, "img")
	err = fsutils.EnsureDir(imgDirectory)
	if err != nil {
		return nil, errors.Wrapf(err, "InitConfig")
	}

	sqlPath := filepath.Join(root, "database.db")

	CONFIG = &Config{
		App: AppConfig{
			Name:       appName,
			Version:    "0.0.1",
			RootDir:    root,
			LogDir:     logDirectory,
			ImageDir:   imgDirectory,
			SqlitePath: sqlPath,
		},
		Database: DatabaseConfig{
			MaxOpenConns:    10,
			MaxIdleConns:    5,
			ConnMaxLifetime: time.Hour * 2,
			PingTimeout:     5 * time.Second,
			Name:            "b3_ux_db",
		},
		HTTP: HTTPConfig{
			Host:             "127.0.0.1",
			Port:             8080,
			WebSocketEnabled: true,
			SecureCookies:    true,
		},
		CORS: CORSConfig{
			Enabled:          true,
			AllowedOrigins:   []string{"http://localhost:4200"},
			AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Origin", "Content-Type", "Authorization"},
			AllowCredentials: true,
		},
	}

	return CONFIG, nil
}
