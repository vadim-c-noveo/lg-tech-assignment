package migrations

import (
  "context"
  "fmt"

  "SwaggerProject/internal/entities"
  "github.com/google/uuid"
  "github.com/uptrace/bun"
)

func init() {
  invitations := []entities.Invitation{
    {ID: uuid.MustParse("ad345bcd-1c0c-4e27-93a8-2948678bc9c4")},
    {ID: uuid.MustParse("b4d3aa6f-b563-425c-9879-b7807e2ff753")},
    {ID: uuid.MustParse("d229d0bc-58ba-4058-95fd-20c1bfb5cbaf")},
  }

  Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
    fmt.Print(" [seeds create] ")
    _, err := db.NewInsert().Model(&invitations).Exec(ctx)
    return err
  }, func(ctx context.Context, db *bun.DB) error {
    fmt.Print(" [seeds destroy] ")
    _, err := db.NewDelete().Model(&invitations).WherePK().Exec(ctx)
    return err
  })
}
