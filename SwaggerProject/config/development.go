package config

import (
  "time"

  "github.com/spf13/viper"
)

type Config struct {
  PostgresDSN            string        `mapstructure:"POSTGRES_DSN"`
  AccessTokenPrivateKey  string        `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
  AccessTokenPublicKey   string        `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
  AccessTokenExpiresIn   time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRES_IN"`
  RefreshTokenPrivateKey string        `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
  RefreshTokenPublicKey  string        `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
  RefreshTokenExpiresIn  time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRES_IN"`
}

func LoadConfig(path string) (config Config, err error) {
  viper.AddConfigPath(path)
  viper.SetConfigType("env")
  viper.SetConfigName("app")
  viper.AutomaticEnv()

  err = viper.ReadInConfig()
  if err != nil {
    return
  }

  err = viper.Unmarshal(&config)
  return
}
