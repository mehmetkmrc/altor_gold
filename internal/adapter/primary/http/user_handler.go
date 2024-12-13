package http

import (
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/mehmetkmrc/ator_gold/internal/converter"

	"github.com/mehmetkmrc/ator_gold/internal/dto"

	"github.com/gofiber/fiber/v3"
)

func (s *server) Login(c fiber.Ctx) error {
	reqBody := new(dto.UserLoginRequest)
	
	body := c.Body()
	if err := json.Unmarshal(body, &reqBody); err != nil {
		return s.errorResponse(c, "error while trying to parse body", err, nil, fiber.StatusBadRequest)
	}
	

	userData, err := s.userService.Login(c.Context(), reqBody.Email, reqBody.Password)
	if err != nil {
		return s.errorResponse(c, "error while trying to login", err, nil, fiber.StatusBadRequest)
	}

	userResponse := converter.GetUserModelToDto(userData.User)
	bearerAccess := "Bearer " + userData.AccessToken
	fmt.Println(userData.AccessToken)
	c.Cookie(&fiber.Cookie{
		Name:     "id",
		Value:    userData.User.UserID,
		Expires:  time.Now().Add(3 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
	})
	c.Cookie(&fiber.Cookie{
		Name:     "name",
		Value:    userData.User.Name + " " + userData.User.Surname,
		Expires:  time.Now().Add(3 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
	})
	c.Cookie(&fiber.Cookie{
		Name:     AccessToken,
		Value:    bearerAccess,
		Expires:  time.Now().Add(time.Hour * 3),
		HTTPOnly: true,
		Secure:   true,
	})
	c.Cookie(&fiber.Cookie{
		Name:     AccessPublic,
		Value:    userData.AccessPublic,
		Expires:  time.Now().Add(time.Hour * 3),
		HTTPOnly: true,
		Secure:   true,
	})

	bearerRefresh := "Bearer " + userData.RefreshToken
	c.Cookie(&fiber.Cookie{
		Name:     RefreshToken,
		Value:    bearerRefresh,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
	})

	c.Cookie(&fiber.Cookie{
		Name:     RefreshPublic,
		Value:    userData.RefreshPublic,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
	})
	

	return s.successResponse(c, userResponse, "Kullanıcı başarıyla giriş yaptı.", fiber.StatusOK)
}

func (s *server) Register(c fiber.Ctx) error {
	reqBody := new(dto.UserRegisterRequest)
	body := c.Body()
	if err := json.Unmarshal(body, &reqBody); err != nil {
		return s.errorResponse(c, "error while trying to parse body", err, nil, fiber.StatusBadRequest)
	}

	
	newUser, err := s.userService.Register(c.Context(), reqBody.Name, reqBody.Surname, reqBody.Email, reqBody.Phone, reqBody.Password)
	if err != nil {
		return s.errorResponse(c, "Error while trying to register user", err, nil, fiber.StatusBadRequest)
	}

	userResponse := converter.GetUserModelToDto(newUser)

	return s.successResponse(c, userResponse, "user registered succesfully", fiber.StatusCreated)
}
