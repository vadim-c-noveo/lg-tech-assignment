package users

import (
  "fmt"

  "SwaggerProject/internal/entities"
  "github.com/uptrace/bun"
  "golang.org/x/net/context"
)

type Repository interface {
  Exists(ctx context.Context, model *entities.User) (bool, error)
  Create(ctx context.Context, model *entities.User) error
}

type repository struct {
  storage *bun.DB
}

func NewRepository(storage *bun.DB) Repository {
  return &repository{storage: storage}
}

func (repo repository) Exists(ctx context.Context, model *entities.User) (bool, error) {
  return repo.storage.NewSelect().Model(model).
    Where("login = ?", model.Login).
    Where("provider = ?", model.Provider).Exists(ctx)
}

func (repo repository) Create(ctx context.Context, model *entities.User) error {
  _, err := repo.storage.NewInsert().Model(model).Exec(ctx)
  if err != nil {
    return fmt.Errorf("users repo: insert: %w", err)
  }

  return nil
}
