package transport

import "github.com/gofiber/fiber/v3"

const AppPostfix = "-bookit"

func Auth() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		//org := ctx.Get("X-Organization")
		//token := ctx.Get("X-Authorization")
		//ctx.S
		return nil
	}
}
