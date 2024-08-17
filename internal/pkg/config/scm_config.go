package config

import (
	"fmt"
	"log"
	"strings"
)

type PostgresConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type SCMConfigS struct {
	Installed      bool           `json:"installed"`
	Port           int            `json:"port"`
	BindAddr       string         `json:"bind_addr"`
	AllowedDomains []string       `json:"allowed_domains"`
	TrustedProxies []string       `json:"trusted_proxies"`
	ReposPath      string         `json:"repos_path"`
	Postgres       PostgresConfig `json:"postgres"`
}

var SCMConfig *SCMConfigS = &SCMConfigS{
	Installed:      false,
	Port:           80,
	BindAddr:       "0.0.0.0",
	AllowedDomains: []string{"example.com"},
	TrustedProxies: []string{""},
	ReposPath:      "{BASE_PATH}/git-data",
	Postgres: PostgresConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "homelab_scm",
		Password: "password",
		Database: "homelab_scm",
	},
}

func init() {
	err := ReadConfig(fmt.Sprintf("%s/configs/scm.json", BasePath), SCMConfig)
	if err != nil {
		log.Fatalf("Failed to read SCM config: %s", err)
	}

	SCMConfig.ReposPath = strings.ReplaceAll(SCMConfig.ReposPath, "{BASE_PATH}", BasePath)
}
