package migrations

import (
  "context"
  "embed"
  "fmt"

  "github.com/uptrace/bun"
  "github.com/uptrace/bun/migrate"
)

var Migrations = migrate.NewMigrations()

//go:embed *.sql
var sqlMigrations embed.FS

func init() {
  if err := Migrations.DiscoverCaller(); err != nil {
    panic(err)
  }
  if err := Migrations.Discover(sqlMigrations); err != nil {
    panic(err)
  }
}

func Migrate(ctx context.Context, db *bun.DB) error {
  migrator := migrate.NewMigrator(db, Migrations)
  err := migrator.Init(ctx)
  if err != nil {
    return err
  }
  if err = migrator.Lock(ctx); err != nil {
    return err
  }
  defer migrator.Unlock(ctx)

  group, err := migrator.Migrate(ctx)
  if err != nil {
    return err
  }
  if group.IsZero() {
    fmt.Printf("there are no new migrations to run (database is up to date)\n")
    return nil
  }
  fmt.Printf("migrated to %s\n", group)

  return nil
}
