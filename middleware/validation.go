package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"httpbst/handler"
)

// validation is simplified
type ValidationMiddleware struct{}

func NewValidationMiddleware() *ValidationMiddleware {
	return &ValidationMiddleware{}
}

func (mw *ValidationMiddleware) Func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var params handler.Params
		if err := ctx.Bind(&params); err != nil {
			return &echo.HTTPError{
				Code:     http.StatusBadRequest,
				Message:  "parsing failed",
				Internal: err,
			}
		}

		if valid := params.Validate(); !valid {
			return &echo.HTTPError{
				Code:    http.StatusUnprocessableEntity,
				Message: "val must be provided",
			}
		}

		ctx.Set("params", params)

		return next(ctx)
	}
}
