package clients

import (
  "database/sql"
  "fmt"

  "github.com/uptrace/bun"
  "github.com/uptrace/bun/dialect/pgdialect"
  "github.com/uptrace/bun/driver/pgdriver"
)

func (c *clients) Postgres() *bun.DB {
  c.postgresOnce.Do(func() {
    dsn := c.config.PostgresDSN
    fmt.Println(dsn)
    sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
    c.postgresClient = bun.NewDB(sqldb, pgdialect.New())
  })
  return c.postgresClient
}
