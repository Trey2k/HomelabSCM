package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"homelabscm.com/scm/internal/pkg/logger"
)

type PostgresConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type SCMConfigS struct {
	DevMode        bool           `json:"dev_mode"`
	BasePath       string         `json:"base_path"`
	Installed      bool           `json:"installed"`
	Port           int            `json:"port"`
	BindAddr       string         `json:"bind_addr"`
	AllowedDomains []string       `json:"allowed_domains"`
	TrustedProxies []string       `json:"trusted_proxies"`
	ReposPath      string         `json:"repos_path"`
	Postgres       PostgresConfig `json:"postgres"`
}

var SCMConfig *SCMConfigS = &SCMConfigS{
	DevMode:        false,
	BasePath:       "/var/opt/homelab-scm",
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
	if os.Getenv("HOMELAB_SCM_BASE_PATH") != "" {
		SCMConfig.BasePath = os.Getenv("HOMELAB_SCM_BASE_PATH")
	}

	if os.Getenv("HOMELAB_SCM_PORT") != "" {
		var err error
		SCMConfig.Port, err = strconv.Atoi(os.Getenv("HOMELAB_SCM_PORT"))
		if err != nil {
			logger.Errorf("Failed to parse HOMELAB_SCM_PORT: %s", err)
		}
	}

	if os.Getenv("HOMELAB_SCM_DEV_MODE") != "" {
		var err error
		SCMConfig.DevMode, err = strconv.ParseBool(os.Getenv("HOMELAB_SCM_DEV_MODE"))
		if err != nil {
			logger.Errorf("Failed to parse HOMELAB_SCM_DEV_MODE: %s", err)
		}
	}

	if os.Getenv("HOMELAB_SCM_BIND_ADDR") != "" {
		SCMConfig.BindAddr = os.Getenv("HOMELAB_SCM_BIND_ADDR")
	}

	err := ReadConfig(fmt.Sprintf("%s/configs/scm.json", SCMConfig.BasePath), SCMConfig)
	if err != nil {
		log.Fatalf("Failed to read SCM config: %s", err)
	}

	SCMConfig.ReposPath = strings.ReplaceAll(SCMConfig.ReposPath, "{BASE_PATH}", SCMConfig.BasePath)
}
