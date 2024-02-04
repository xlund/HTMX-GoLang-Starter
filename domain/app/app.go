package app

import (
	"fmt"
	"go-starter/router"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	server ServerConfig `yaml:"server"`
}

func (c *AppConfig) UnmarshalYAML(data []byte) error {
	var cfg struct {
		Server struct {
			Host string `yaml:"host"`
			Port int    `yaml:"port"`
		} `yaml:"server"`
	}
	err := yaml.Unmarshal(data, &cfg)
	if err != nil {
		return err
	}
	c.server.host = cfg.Server.Host
	c.server.port = cfg.Server.Port
	return nil
}

type ServerConfig struct {
	host string `yaml:"host"`
	port int    `yaml:"port"`
}

func (c *AppConfig) ServerAddress() string {
	return fmt.Sprintf("%s:%d", c.server.host, c.server.port)
}

type App struct {
	Router *chi.Mux
	Server *http.Server
}

func ReadConfig() (AppConfig, error) {
	f, err := os.ReadFile("configuration/base.yaml")
	println("Reading configuration...")
	if err != nil {
		return AppConfig{}, err
	}

	var c AppConfig
	err = c.UnmarshalYAML(f)
	if err != nil {
		return AppConfig{}, err
	}
	return c, nil
}

func New() (*App, error) {
	config, err := ReadConfig()
	if err != nil {
		return nil, err
	}
	r := router.New()
	s := &http.Server{
		Addr:    config.ServerAddress(),
		Handler: r,
	}
	println("Server created...")
	return &App{Router: r, Server: s}, nil
}
