package transport

import (
	"github.com/gofiber/fiber/v3"

	"bookit/pkg/errors"
	"bookit/pkg/iam"
)

func Signin(ctx fiber.Ctx) error {
	code, ok := ctx.Queries()["code"]
	if !ok {
		return errors.New("code not found", errors.WithHttpStatus(fiber.StatusForbidden))
	}
	state, ok := ctx.Queries()["state"]
	if !ok {
		return errors.New("state not found", errors.WithHttpStatus(fiber.StatusForbidden))
	}
	token, err := iam.Gist().GetOAuthToken(code, state)
	if err != nil {
		return errors.Wrap(err, errors.WithHttpStatus(fiber.StatusForbidden))
	}
	return ctx.JSON(token)
}

func User(ctx fiber.Ctx) error {
	token := getAccessToken(ctx)
	claims, err := iam.Gist().ParseJwtToken(token)
	if err != nil {
		return errors.Wrap(err, errors.WithHttpStatus(fiber.StatusForbidden))
	}
	return ctx.JSON(claims.User)
}
