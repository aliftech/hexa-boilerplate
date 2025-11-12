package response

import (
	"net/http"
	"strings"

	"hexa-fiber-gorm/internal/i18n"

	"github.com/gofiber/fiber/v2"
)

// Success returns a localized success response
func Success(c *fiber.Ctx, httpStatus int, key string, data interface{}) error {
	lang := getLangFromCtx(c)
	message := i18n.Get(lang, key)
	return c.Status(httpStatus).JSON(fiber.Map{
		"status":  "success",
		"message": message,
		"data":    data,
	})
}

// Error returns a localized error response
func Error(c *fiber.Ctx, httpStatus int, key string) error {
	lang := getLangFromCtx(c)
	message := i18n.Get(lang, key)
	return c.Status(httpStatus).JSON(fiber.Map{
		"status":  "error",
		"message": message,
		"error":   http.StatusText(httpStatus),
	})
}

// getLangFromCtx extracts and normalizes language from Accept-Language header
func getLangFromCtx(c *fiber.Ctx) string {
	lang := c.Get("Accept-Language", "en")
	if len(lang) == 0 {
		return "en"
	}
	if len(lang) >= 2 {
		lang = strings.ToLower(lang[:2])
	}
	// Only allow known languages
	if lang == "id" || lang == "en" {
		return lang
	}
	return "en"
}
