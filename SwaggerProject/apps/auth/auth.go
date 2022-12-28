package auth

import (
  "SwaggerProject/config"
  "SwaggerProject/internal/invitations"
  "SwaggerProject/internal/users"
  "github.com/gin-gonic/gin"
)

type App interface {
  Register(c *gin.Context)
}

func NewApp(
  usersRepository users.Repository,
  invitationsRepository invitations.Repository,
  cfg *config.Config,
) App {
  return &app{usersRepository, invitationsRepository, cfg}
}

type app struct {
  UsersRepository       users.Repository
  InvitationsRepository invitations.Repository
  // TODO: should be an interface
  config *config.Config
}
