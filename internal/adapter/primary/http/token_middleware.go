package http

import (
	"github.com/gofiber/fiber/v3"
	"strings"
)

func (s *server) isValidToken(c fiber.Ctx) bool {
	token := c.Cookies(AccessToken)
	if token == "" {
		return false
	}

	fields := strings.Fields(token)
	if len(fields) != 2 || fields[0] != AuthType {
		return false
	}

	return true
}

func (s *server) isValidPublicKey(c fiber.Ctx) bool {
	publicKey := c.Cookies(AccessPublic)
	return publicKey != ""
}

func (s *server) getAccessToken(c fiber.Ctx) string {
	fields := strings.Fields(c.Cookies(AccessToken))
	return fields[1]
}

func (s *server) getAccessPublicKey(c fiber.Ctx) string {
	return c.Cookies(AccessPublic)
}
