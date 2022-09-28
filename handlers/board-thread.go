package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"tewi/config"
	"tewi/db"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/uptrace/bun"
)

// BoardThread loads up a specific thread given the thread number
func BoardThread(pg *bun.DB, conf config.Config) func(echo.Context) error {
	return func(c echo.Context) error {
		board := c.Param("board")

		if !lo.Contains(conf.Boards, board) {
			return c.String(http.StatusBadRequest, fmt.Sprintf("Board \"%s\" not available"))
		}

		threadNumberS := c.Param("thread_number")
		threadNumber, err := strconv.ParseInt(threadNumberS, 10, 64)

		if err != nil {
			return c.String(http.StatusBadRequest, "Error processing thread number parameter")
		}

		thread := make([]db.Post, 0, 10)
		err = pg.NewSelect().
			Model(&thread).
			Where("board = ?", board).
			Where("thread_number = ?", threadNumber).
			Where("NOT hidden").
			Order("post_number DESC").
			Scan(context.Background())

		if err != nil {
			return c.String(http.StatusInternalServerError, "Internal server error")
		}

		//If the thread is empty of the thread OP has been hidden (hence it's not the first element in the array)
		//Ideally if a thread is hidden all posts in it should also be hidden in the DB, this is just to make sure.
		if len(thread) == 0 || !thread[0].Op {
			return c.String(http.StatusNotFound, "Thread not found")
		}

		return c.JSON(http.StatusOK, thread)
	}
}
