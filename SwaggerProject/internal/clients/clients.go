package clients

import (
  "sync"

  "SwaggerProject/config"
  "github.com/uptrace/bun"
)

type Clients interface {
  Postgres() *bun.DB
}

type clients struct {
  config         *config.Config
  postgresOnce   sync.Once
  postgresClient *bun.DB
}

func New(config *config.Config) Clients {
  return &clients{
    config: config,
  }
}
