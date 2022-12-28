package migrations

import (
  "context"
  "fmt"

  "SwaggerProject/internal/entities"
  "github.com/uptrace/bun"
)

func init() {
  Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
    fmt.Print(" [up migration] ")
    _, err := db.NewCreateTable().IfNotExists().Model((*entities.Invitation)(nil)).Exec(ctx)
    return err
  }, func(ctx context.Context, db *bun.DB) error {
    fmt.Print(" [down migration] ")
    _, err := db.NewDropTable().Model((*entities.Invitation)(nil)).IfExists().Exec(ctx)
    return err
  })
}
