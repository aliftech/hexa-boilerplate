package web

import (
	"strconv"

	ports "hexa-fiber-gorm/internal/modules/user/port"
	"hexa-fiber-gorm/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	usecase ports.UserUsecase
}

func NewHandler(usecase ports.UserUsecase) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	req := struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}{}

	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, 400, "http.bad_request")
	}

	user, err := h.usecase.CreateUser(req.Name, req.Email)
	if err != nil {
		// TODO: inspect err type for better key (e.g., "user.email_exists")
		return response.Error(c, 409, "user.email_exists")
	}

	return response.Success(c, 201, "user.created", user)
}

func (h *Handler) GetUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return response.Error(c, 400, "http.bad_request")
	}

	user, err := h.usecase.GetUserByID(uint(id))
	if err != nil {
		return response.Error(c, 404, "user.not_found")
	}

	return response.Success(c, 200, "http.ok", user)
}

func (h *Handler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.usecase.GetAllUsers()
	if err != nil {
		return response.Error(c, 500, "http.internal_error")
	}

	return response.Success(c, 200, "http.ok", users)
}
