package handler

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"zsxyww.com/scheduler/database"
	"zsxyww.com/scheduler/model"
)

var cxt = context.Background()

func HandlerTest(i echo.Context) error {

	switch i.QueryParam("acts") {
	case "test":
		if err := db.PGX.Ping(cxt); err != nil {
			return i.String(500, err.Error())
		}
		return i.String(200, "ok")
	case "select":
		a, _ := db.PGX.Query(cxt, "select * from members")
		result, err := pgx.CollectRows(a, pgx.RowToStructByName[model.Member])
		if err != nil {
			return i.String(500, err.Error())
		}
		return i.JSON(200, result)
	default:
		return i.String(200, "give arguments please")
	}
}
