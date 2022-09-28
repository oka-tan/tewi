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

// BoardPost looks up an individual post
func BoardPost(pg *bun.DB, conf config.Config) func(echo.Context) error {
	return func(c echo.Context) error {
		board := c.Param("board")

		if !lo.Contains(conf.Boards, board) {
			return c.String(http.StatusBadRequest, fmt.Sprintf("Board \"%s\" not available", board))
		}

		postNumberS := c.Param("post_number")
		postNumber, err := strconv.ParseInt(postNumberS, 10, 64)

		if err != nil {
			return c.String(http.StatusBadRequest, "Error processing post number parameter")
		}

		var post db.Post

		err = pg.NewSelect().
			Model(&post).
			Where("board = ?", board).
			Where("post_number = ?", postNumber).
			Where("NOT hidden").
			Scan(context.Background())

		if err != nil {
			return c.String(http.StatusNotFound, "Post not found")
		}

		return c.JSON(http.StatusOK, post)
	}
}
