package entities

import (
  "github.com/google/uuid"
  "github.com/uptrace/bun"
)

type User struct {
  bun.BaseModel `bun:"table:users"`

  ID       int64 `bun:",pk,autoincrement"`
  Login    string
  Password string
  Provider int
  UUID     uuid.UUID
}
