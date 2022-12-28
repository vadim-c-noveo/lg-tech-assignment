package app

import (
  "context"
  "fmt"
  "net/http"
  "time"

  "SwaggerProject/apps/auth"
  "SwaggerProject/config"
  "SwaggerProject/internal/clients"
  "SwaggerProject/internal/invitations"
  "SwaggerProject/internal/users"
  "SwaggerProject/migrations"
  "github.com/gin-gonic/gin"
)

type Server struct {
  *http.Server
  repositories
}

type repositories struct {
  users       users.Repository
  invitations invitations.Repository
}

func NewServer(router *gin.Engine) *Server {
  return &Server{
    Server: &http.Server{
      Addr:           ":8080",
      Handler:        router,
      ReadTimeout:    10 * time.Second,
      WriteTimeout:   10 * time.Second,
      MaxHeaderBytes: 1 << 20,
    },
  }
}

func Start(ctx context.Context) error {
  cfg, err := config.LoadConfig(".")
  if err != nil {
    return err
  }

  cl := clients.New(&cfg)
  repositories := repositories{
    users:       users.NewRepository(cl.Postgres()),
    invitations: invitations.NewRepository(cl.Postgres()),
  }

  err = migrations.Migrate(ctx, cl.Postgres())
  if err != nil {
    return fmt.Errorf("app start: %w", err)
  }

  authApp := auth.NewApp(repositories.users, repositories.invitations, &cfg)
  router := gin.Default()
  router.POST("/auth/register", authApp.Register)

  s := NewServer(router)

  return s.ListenAndServe()
}
