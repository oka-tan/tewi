package main

import (
	"database/sql"
	"fmt"
	"tewi/config"
	"tewi/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	gommon_log "github.com/labstack/gommon/log"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func main() {
	conf := config.LoadConfig()

	sqlpg := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(conf.PostgresConfig.ConnectionString)))
	pg := bun.NewDB(sqlpg, pgdialect.New())

	e := echo.New()

	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10,
		LogLevel:  gommon_log.ERROR,
	}))

	e.Use(middleware.Gzip())

	e.GET("/:board", handlers.Board(pg, conf))
	e.GET("/:board/post/:post_number", handlers.BoardPost(pg, conf))
	e.GET("/:board/thread/:thread_number", handlers.BoardThread(pg, conf))
	e.GET("/:board/view-same/:media_4chan_hash", handlers.BoardViewSame(pg, conf))

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", conf.Port)))
}
