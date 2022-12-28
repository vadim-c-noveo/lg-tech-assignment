package auth

import (
  "net/http"
  "strconv"

  "SwaggerProject/internal/entities"
  "SwaggerProject/internal/jwt"
  "github.com/gin-gonic/gin"
  "github.com/go-playground/validator/v10"
  "github.com/google/uuid"
)

type registerRequest struct {
  Login           string `json:"login" validate:"required,alphanum,gte=6,lte=255"`
  Password        string `json:"password" validate:"required,gte=8,lte=255"`
  UUID            string `json:"UUID" validate:"required,uuid_rfc4122"`
  PrivacyAccepted bool   `json:"privacyAccepted" validate:"required"`
}

type RegisterRequestValidator struct{}

func NewRegisterRequestValidator() *RegisterRequestValidator {
  return &RegisterRequestValidator{}
}

var validate *validator.Validate

func (v *RegisterRequestValidator) Validate(request *registerRequest) error {
  validate = validator.New()
  err := validate.Struct(request)
  if err != nil {
    return err
  }
  return nil
}

func (app *app) Register(c *gin.Context) {
  var request registerRequest
  if err := c.ShouldBindJSON(&request); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  requestValidator := NewRegisterRequestValidator()
  err := requestValidator.Validate(&request)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  parsedUUID, err := uuid.Parse(request.UUID)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invitation not found"})
    return
  }

  provider, err := strconv.Atoi(c.GetHeader("X-Provider"))
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  invitationExists, err := app.InvitationsRepository.Exists(c.Request.Context(), parsedUUID)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }
  if !invitationExists {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Invitation not found"})
    return
  }

  user := entities.User{
    Login:    request.Login,
    Password: request.Password,
    Provider: provider,
    UUID:     parsedUUID,
  }
  userExists, err := app.UsersRepository.Exists(c.Request.Context(), &user)
  if userExists {
    c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
    return
  }

  err = app.UsersRepository.Create(c.Request.Context(), &user)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  accessToken, err := jwt.CreateToken(user.ID, app.config.AccessTokenExpiresIn, app.config.AccessTokenPrivateKey)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
  }

  refreshToken, err := jwt.CreateToken(user.ID, app.config.RefreshTokenExpiresIn, app.config.RefreshTokenPrivateKey)
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "login":        user.Login,
    "refreshToken": refreshToken,
    "accessToken":  accessToken,
    "provider":     provider,
    "locale":       "fr_FR",
  })
}
