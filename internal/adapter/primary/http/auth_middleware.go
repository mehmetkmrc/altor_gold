package http

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/mehmetkmrc/ator_gold/internal/converter"
	"github.com/mehmetkmrc/ator_gold/internal/core/domain/valueobject"
)

const (
	AuthHeader    = "Authorization"
	AccessToken   = "access_token"
	AccessPublic  = "access_public"
	RefreshToken  = "refresh_token"
	RefreshPublic = "refresh_public"
	UserDetail    = "UserDetail"
	AuthType      = "Bearer"
	AuthPayload   = "Payload"
)

func (s *server) IsAuthorized(c fiber.Ctx) error {

	fmt.Println("IsAuthorized middleware çalıştı")
	if !s.isValidToken(c) {
		return s.redirectToLogin(c, fiber.StatusUnauthorized, "authorization header is not provided or invalid")
	}

	if !s.isValidPublicKey(c) {
		return s.redirectToLogin(c, fiber.StatusUnauthorized, "public key is not provided")
	}

	token := s.getAccessToken(c)
	publicKey := s.getAccessPublicKey(c)

	payload, err := s.tokenService.DecodeToken(token, publicKey)
	if err != nil {
		return s.redirectToLogin(c, fiber.StatusUnauthorized, "invalid access token")
	}

	c.Locals(AuthPayload, payload)
	return c.Next()
}

func (s *server) GetUserDetail(c fiber.Ctx) error {
	payload := c.Locals(AuthPayload).(*valueobject.Payload)

	userAggregate, err := s.userService.GetUserByID(c.Context(), payload.ID)
	if err != nil {
		return s.errorResponse(c, "error while trying to get user detail", err, nil, fiber.StatusBadRequest)
	}

	userResponse := converter.GetUserModelToDto(userAggregate)
	c.Locals(UserDetail, userResponse)
	return c.Next()
}

func (s *server) redirectToLogin(c fiber.Ctx, statusCode int, message string) error {
	c.Status(statusCode).JSON(fiber.Map{
		"error": message,
	})
	return c.Redirect().To("/")
}


