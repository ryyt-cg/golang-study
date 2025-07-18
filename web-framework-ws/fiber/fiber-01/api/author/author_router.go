package author

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type Router struct {
	authorService Servicer
}

// NewRouter creates a new Router
func NewRouter(authorService Servicer) *Router {
	return &Router{
		authorService,
	}
}

// Register registers the router to the gin engine
func (authorRouter *Router) Register(router fiber.Router) {
	router.Get(":id", authorRouter.authorByID)
}

// appInfo	Show app info
func (authorRouter *Router) authorByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		log.Error().Err(err).Str("id", c.Params("id")).Msg("Invalid author ID")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid author ID",
		})
	}

	log.Debug().Int("ID", id).Msg("Fetching author by ID")
	result, err := authorRouter.authorService.getAuthorByID(id)
	if err != nil {
		log.Error().Err(err).Msg("Failed to fetch author by ID")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch author",
		})
	}

	return c.JSON(result)
}
