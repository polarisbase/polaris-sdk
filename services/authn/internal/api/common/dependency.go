package common

import "github.com/gofiber/fiber/v2"

type Dependencies struct {
	FiberRouter fiber.Router
}

func NewDependencies(fiberRouter fiber.Router) *Dependencies {

	d := &Dependencies{}

	d.FiberRouter = fiberRouter

	return d

}
