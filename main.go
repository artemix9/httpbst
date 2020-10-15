package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"httpbst/config"
	"httpbst/handler"
	"httpbst/middleware"
	"httpbst/tree"
)

func main() {
	conf, err := config.Load("config.json")
	if err != nil {
		panic(err)
	}

	if err := conf.Validate(); err != nil {
		panic(err)
	}

	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	var t tree.Tree = tree.NewPureTree(conf.Values)
	t = tree.NewLoggableTree(t, e.Logger)
	t = tree.NewSafeTree(t)

	h := handler.NewHandler(t)

	mw := middleware.NewValidationMiddleware()

	e.Use(mw.Func)

	e.GET("/search", h.Search)
	e.POST("/insert", h.Insert)
	e.DELETE("/delete", h.Delete)

	if err := e.Start(":8000"); err != nil {
		panic(err)
	}
}
