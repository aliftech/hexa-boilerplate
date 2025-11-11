// internal/modules/user/adapter/web/handler.go
package web

import (
	"hexa-fiber-gorm/internal/modules/user/domain"
	ports "hexa-fiber-gorm/internal/modules/user/port"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	repo ports.UserRepository
}

func NewHandler(repo ports.UserRepository) *Handler {
	return &Handler{repo: repo}
}

// CreateUser handles the HTTP request to create a user
func (h *Handler) CreateUser(c *fiber.Ctx) error {
	user := new(domain.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid payload"})
	}
	if err := h.repo.Create(user); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "failed to create user"})
	}
	return c.Status(201).JSON(user)
}

// GetUser handles fetching a user by ID
func (h *Handler) GetUser(c *fiber.Ctx) error {
	// id := c.Params("id")
	// TODO: parse id to uint (use strconv)
	// For now, simplified
	user, err := h.repo.FindByID(1)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "user not found"})
	}
	return c.JSON(user)
}

// GetAllUsers returns all users
func (h *Handler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.repo.FindAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "database error"})
	}
	return c.JSON(users)
}
