package transport

import (
	"github.com/gofiber/fiber/v3"

	"bookit/pkg/config"
)

// Version	Получить версию
//
//	@summary	Получить версию приложения
//	@tags		app
//	@id			version
//	@success	200
//	@failure	400
//	@failure	401
//	@failure	403
//	@failure	500
//	@router		/version [get]
func Version(ctx fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).SendString(config.Gist().App.Version)
}

// Health	Проверка работоспособности
//
//	@summary	Проверка работоспособности
//	@tags		app
//	@id			health
//	@success	200
//	@failure	400
//	@failure	401
//	@failure	403
//	@failure	500
//	@router		/healthz [get]
func Health(ctx fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).SendString("OK")
}

// Panic	Проверка panic-recovery
//
//	@summary	Проверка panic-recovery
//	@tags		app
//	@id			panic
//	@success	200
//	@failure	400
//	@failure	401
//	@failure	403
//	@failure	500
//	@router		/panic [get]
func Panic(_ fiber.Ctx) error {
	panic("panic")
}
