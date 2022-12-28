package entities

import (
  "github.com/google/uuid"
  "github.com/uptrace/bun"
)

type Invitation struct {
  bun.BaseModel `bun:"table:invitations"`

  ID uuid.UUID `bun:",pk,unique,type:uuid,default:uuid_generate_v4()"`
}
