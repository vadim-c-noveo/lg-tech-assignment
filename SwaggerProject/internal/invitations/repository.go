package invitations

import (
  "context"
  "fmt"

  "SwaggerProject/internal/entities"
  "github.com/google/uuid"
  "github.com/uptrace/bun"
)

// TODO: merge with users/repository and abstract it if there would be time left

type Repository interface {
  Exists(ctx context.Context, id uuid.UUID) (bool, error)
  Create(ctx context.Context, model *entities.Invitation) error
}

type repository struct {
  storage *bun.DB
}

func NewRepository(storage *bun.DB) Repository {
  return &repository{storage: storage}
}

func (repo repository) Exists(ctx context.Context, id uuid.UUID) (bool, error) {
  invitation := new(entities.Invitation)
  return repo.storage.NewSelect().Model(invitation).Where("id = ?", id).Exists(ctx)
}

func (repo repository) Create(ctx context.Context, model *entities.Invitation) error {
  _, err := repo.storage.NewInsert().Model(model).Exec(ctx)
  if err != nil {
    return fmt.Errorf("invitations repo: insert: %w", err)
  }

  return nil
}
