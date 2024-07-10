package transport

import (
	"strings"

	"github.com/gofiber/fiber/v3"
	"gopkg.in/resty.v1"

	"bookit/pkg/config"
	"bookit/pkg/errors"
	"bookit/pkg/iam"
)

var noAuthPaths = map[string]struct{}{
	"/v1/version": {},
	"/v1/healthz": {},
	"/v1/signin":  {},
}

func Auth() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		if !config.Gist().App.AuthEnabled ||
			isNoAuthPath(ctx) {
			return ctx.Next()
		}
		token := getAccessToken(ctx)
		client := resty.New()
		validationResult := make(map[string]interface{})
		_, err := client.R().
			SetBasicAuth(config.Gist().Iam.ClientId, config.Gist().Iam.ClientSecret).
			SetHeader("Content-Type", "application/x-www-form-urlencoded").
			SetQueryParams(map[string]string{
				"token":           token,
				"token_type_hint": "access_token",
			}).
			SetResult(&validationResult).
			Post(config.Gist().Iam.Url + "/api/login/oauth/introspect")
		if err != nil {
			return errors.Wrap(err, errors.WithHttpStatus(fiber.StatusForbidden))
		}
		if validationResult["active"].(bool) == false {
			return errors.New("invalid token", errors.WithHttpStatus(fiber.StatusForbidden))
		}
		claims, err := iam.Gist().ParseJwtToken(token)
		if err != nil {
			return errors.Wrap(err, errors.WithHttpStatus(fiber.StatusForbidden))
		}
		ctx.Locals("user", claims.User)
		return ctx.Next()
	}
}

func getAccessToken(ctx fiber.Ctx) string {
	return strings.Replace(ctx.Get("Authorization"), "Bearer ", "", 1)
}

func isNoAuthPath(ctx fiber.Ctx) bool {
	path := ctx.Path()
	_, ok := noAuthPaths[path]
	return ok
}
