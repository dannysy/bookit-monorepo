package transport

import (
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/ztrue/tracerr"

	"bookit/pkg/iam"
)

func Signin(ctx fiber.Ctx) error {
	code, ok := ctx.Queries()["code"]
	if !ok {
		return fiber.ErrForbidden
	}
	state, ok := ctx.Queries()["state"]
	if !ok {
		return fiber.ErrForbidden
	}
	token, err := iam.Gist().GetOAuthToken(code, state)
	if err != nil {
		return tracerr.Wrap(err)
	}
	return ctx.JSON(token)
}

func User(ctx fiber.Ctx) error {
	token := strings.Replace(ctx.Get("Authorization"), "Bearer ", "", 1)
	claims, err := iam.Gist().ParseJwtToken(token)
	if err != nil {
		return tracerr.Wrap(err)
	}
	return ctx.JSON(claims.User)
}
