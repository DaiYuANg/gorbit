package jwt

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/oops"
	"go.uber.org/fx"
)

func NewJwtModule(opts ...Option) fx.Option {
	options := defaultJwtOptions()
	for _, o := range opts {
		o(options)
	}

	return fx.Module("jwt_module",
		fx.Invoke(func(app *fiber.App) error {
			if options.PrivateKey == nil {
				return oops.New("JWT private key is required")
			}

			app.Use(options.PathPrefix, jwtware.New(jwtware.Config{
				SigningKey: jwtware.SigningKey{
					JWTAlg: options.SigningAlg,
					Key:    options.PrivateKey.Public(),
				},
			}))
			return nil
		}),
	)
}
