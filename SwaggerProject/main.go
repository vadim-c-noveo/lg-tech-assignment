package main

import (
  "context"
  "log"

  "SwaggerProject/internal/app"
)

func main() {
  ctx := context.Background()
  err := app.Start(ctx)
  if err != nil {
    log.Fatal(err)
  }
}
