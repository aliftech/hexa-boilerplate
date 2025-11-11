package web

import (
	"hexa-fiber-gorm/internal/modules/user/adapter/web/dto"
	ports "hexa-fiber-gorm/internal/modules/user/port"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	usecase ports.UserUsecase // ✅ depends on USECASE, not repo
}

func NewHandler(usecase ports.UserUsecase) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) CreateUser(c *fiber.Ctx) error {
	req := new(dto.CreateUserRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid payload"})
	}

	user, err := h.usecase.CreateUser(req.Name, req.Email)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(user)
}

func (h *Handler) GetUser(c *fiber.Ctx) error {
	// id := c.Params("id")
	// TODO: parse id to uint
	user, err := h.usecase.GetUserByID(1)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "user not found"})
	}
	return c.JSON(user)
}

func (h *Handler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.usecase.GetAllUsers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to fetch users"})
	}
	return c.JSON(users)
}
