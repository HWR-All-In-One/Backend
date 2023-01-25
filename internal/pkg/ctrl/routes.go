package ctrl

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func (a *App) AddTimetableRoutes() {
	a.PB.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// or you can also use the shorter e.Router.GET("/articles/:slug", handler, middlewares...)
		e.Router.AddRoute(echo.Route{
			Method:  http.MethodGet,
			Path:    "/hwr/timetable/list",
			Handler: a.Views.Timetable.List,
			Middlewares: []echo.MiddlewareFunc{
				apis.ActivityLogger(a.PB),
			},
		})
		return nil
	})

}
