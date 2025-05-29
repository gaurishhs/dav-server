package config_test

import (
	"errors"
	"os"
	"reflect"
	"testing"

	"github.com/gaurishhs/dav-server/internal/config"
)

var testConfigPath = "/config_test.toml"

var testConfig *config.Config = &config.Config{
	Server: config.ServerConfig{
		Addr: ":8080",
	},
}

func ValidateConfig(config *config.Config) error {
	if !reflect.DeepEqual(testConfig, config) {
		return errors.New("validation failed")
	}
	return nil
}

func TestLoadConfig(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	config, err := config.LoadConfig(cwd + testConfigPath)
	if err != nil {
		t.Fatal(err)
	}
	if config == nil {
		t.Fatal("config is nil")
	}
	if err := ValidateConfig(config); err != nil {
		t.Fatal(err)
	}
}
