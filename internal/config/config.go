package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

var (
  configDir = "gnews"
  configFile = "config.toml"
  defaultDbName = "gnews"
  defaultDbSystem = "bbolt"
)

type Config struct {
  Database Database `toml:"database"`
}

type Database struct {
  Name string `toml:"name"`
  System string `toml:"system"`
}

func GetConfig() (Config, error) {
  dir, err := os.UserConfigDir()
  if err != nil {
    log.Fatal(err)
  }

  configFilePath := filepath.Join(dir, configDir, configFile)
  _, err = os.Stat(configFilePath)
  if os.IsNotExist(err) {
    return Config{}, err
  }

  data, err := os.ReadFile(configFilePath)
  if err != nil {
    log.Fatal(err)
  }

  var cfg Config
  err = toml.Unmarshal([]byte(data), &cfg)
  if err != nil {
    log.Fatal(err)
  }

  return cfg, nil
}

func InitConfig() Config {
  dir, err := os.UserConfigDir()
  if err != nil {
    log.Fatal(err)
  }

  configDirPath := filepath.Join(dir, configDir)
  err = os.MkdirAll(configDirPath, 0750)
  if err != nil {
    log.Fatal(err)
  }

  config := Config{Database: Database{Name: defaultDbName, System: defaultDbSystem}}
  configFilePath := filepath.Join(dir, configDir, configFile)
  c, err := toml.Marshal(config)
  if err != nil {
    log.Fatal(err)
  }

  err = os.WriteFile(configFilePath, []byte(c), 0660)
  if err != nil {
    log.Fatal(err)
  }

  return config
}

func ConfigDir() string {
  dir, err := os.UserConfigDir()
  if err != nil {
    log.Fatal(err)
  }
  
  configDirPath := filepath.Join(dir, configDir)

  return configDirPath
}
