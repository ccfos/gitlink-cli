package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const (
	DefaultBaseURL = "https://www.gitlink.org.cn/api"
	DefaultFormat  = "table"
)

type Config struct {
	BaseURL string `yaml:"base_url"`
	Format  string `yaml:"default_format"`
	Editor  string `yaml:"editor,omitempty"`
	Pager   string `yaml:"pager,omitempty"`
}

func DefaultConfig() *Config {
	return &Config{
		BaseURL: DefaultBaseURL,
		Format:  DefaultFormat,
	}
}

func ConfigDir() string {
	if dir := os.Getenv("GITLINK_CONFIG_DIR"); dir != "" {
		return dir
	}
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "gitlink-cli")
}

func ConfigPath() string {
	return filepath.Join(ConfigDir(), "config.yaml")
}

func Load() (*Config, error) {
	cfg := DefaultConfig()
	data, err := os.ReadFile(ConfigPath())
	if err != nil {
		if os.IsNotExist(err) {
			return cfg, nil
		}
		return nil, err
	}
	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}
	if cfg.BaseURL == "" {
		cfg.BaseURL = DefaultBaseURL
	}
	if cfg.Format == "" {
		cfg.Format = DefaultFormat
	}
	return cfg, nil
}

func Save(cfg *Config) error {
	dir := ConfigDir()
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(ConfigPath(), data, 0600)
}

func Get(key string) (string, error) {
	cfg, err := Load()
	if err != nil {
		return "", err
	}
	switch key {
	case "base_url":
		return cfg.BaseURL, nil
	case "default_format":
		return cfg.Format, nil
	case "editor":
		return cfg.Editor, nil
	case "pager":
		return cfg.Pager, nil
	default:
		return "", nil
	}
}

func Set(key, value string) error {
	cfg, err := Load()
	if err != nil {
		return err
	}
	switch key {
	case "base_url":
		cfg.BaseURL = value
	case "default_format":
		cfg.Format = value
	case "editor":
		cfg.Editor = value
	case "pager":
		cfg.Pager = value
	}
	return Save(cfg)
}
