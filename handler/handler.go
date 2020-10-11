package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"httpbst/tree"
)

type Handler struct {
	t tree.Tree
}

func NewHandler(t tree.Tree) *Handler {
	return &Handler{t: t}
}

type Params struct {
	Val *int
}

func (params Params) Validate() bool {
	return params.Val != nil
}

func (h *Handler) Search(ctx echo.Context) error {
	params := ctx.Get("params").(Params)

	if err := ctx.JSON(http.StatusOK, h.t.Search(*params.Val)); err != nil {
		return fmt.Errorf("handler: search: send response: %s", err)
	}

	return nil
}

func (h *Handler) Insert(ctx echo.Context) error {
	params := ctx.Get("params").(Params)

	if err := ctx.JSON(http.StatusOK, h.t.Insert(*params.Val)); err != nil {
		return fmt.Errorf("handler: insert: send response: %s", err)
	}

	return nil
}

func (h *Handler) Delete(ctx echo.Context) error {
	params := ctx.Get("params").(Params)

	if err := ctx.JSON(http.StatusOK, h.t.Delete(*params.Val)); err != nil {
		return fmt.Errorf("handler: delete: send response: %s", err)
	}

	return nil
}
