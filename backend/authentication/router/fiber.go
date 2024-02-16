package router

import (
	"mtrain-main/usecases"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type FiberContext struct {
	*fiber.Ctx
}

func NewFiberContext(c *fiber.Ctx) *FiberContext {
	return &FiberContext{Ctx: c}
}

func (c *FiberContext) Bind(v interface{}) error {
	return c.Ctx.BodyParser(v)
}

func (c *FiberContext) JSON(statuscode int, v interface{}) {
	c.Ctx.Status(statuscode).JSON(v)
}

type FiberRouter struct {
	*fiber.App
}

func NewFiberRouter() *FiberRouter {
	r := fiber.New()
	r.Use(cors.New())
	r.Use(logger.New())

	return &FiberRouter{r}
}

type HandlerFunc func(usecases.Context)

func (r *FiberRouter) GET(relativePath string, handler HandlerFunc) {
	r.App.Get(relativePath, func(c *fiber.Ctx) error {
		handler(NewFiberContext(c))
		return nil
	})
}
