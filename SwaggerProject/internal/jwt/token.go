package jwt

import (
  "encoding/base64"
  "fmt"
  "time"

  "github.com/golang-jwt/jwt/v4"
)

func CreateToken(payload interface{}, ttl time.Duration, privateKey string) (string, error) {
  decodedPrivateKey, err := base64.StdEncoding.DecodeString(privateKey)
  if err != nil {
    return "", fmt.Errorf("create token: key decode: %w", err)
  }
  key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)
  if err != nil {
    return "", fmt.Errorf("create token: key parse: %w", err)
  }

  now := time.Now().UTC()

  claims := make(jwt.MapClaims)
  claims["sub"] = payload
  claims["exp"] = now.Add(ttl).Unix()
  claims["iat"] = now.Unix()

  token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
  if err != nil {
    return "", fmt.Errorf("create token: sign token: %w", err)
  }

  return token, nil
}
