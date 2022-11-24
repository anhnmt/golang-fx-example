package main

import (
	"net/http"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Provide(
			NewHTTPServer,
			NewServeMux,
			fx.Annotate(
				NewEchoHandler,
				fx.As(new(Route)),
			),
			zap.NewExample,
		),
		fx.Invoke(func(*http.Server) {}),
	).Run()

}
