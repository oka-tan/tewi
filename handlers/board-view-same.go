package handlers

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"
	"tewi/config"
	"tewi/db"

	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"github.com/uptrace/bun"
)

// BoardViewSame performs a view-same search for posts with a specific md5 hash
func BoardViewSame(pg *bun.DB, conf config.Config) func(echo.Context) error {
	return func(c echo.Context) error {
		board := c.Param("board")

		boardExists := lo.Contains(conf.Boards, board)

		if !boardExists {
			return c.String(http.StatusBadRequest, fmt.Sprintf("Board \"%s\" not found", board))
		}

		media4chanHashString := c.Param("media_4chan_hash")
		media4chanHash, err := base64.URLEncoding.DecodeString(media4chanHashString)

		if err != nil {
			return c.String(http.StatusBadRequest, fmt.Sprintf("Media hash not properly base64 encoded."))
		}

		var posts []db.Post

		q := pg.NewSelect().
			Model(&posts).
			Where("board = ?", board).
			Where("media_4chan_hash = ?", media4chanHash).
			Where("NOT hidden")

		keyset, kerr := strconv.ParseInt(c.QueryParam("keyset"), 10, 64)
		rkeyset, rkerr := strconv.ParseInt(c.QueryParam("rkeyset"), 10, 64)

		if kerr == nil {
			q.Where("post_number < ?", keyset)
		} else if rkerr == nil {
			q.Where("post_number > ?", rkeyset)
		}

		err = q.Order("post_number DESC").
			Limit(24).
			Scan(context.Background())

		if err != nil {
			return c.String(http.StatusInternalServerError, "Internal server error")
		}

		return c.JSON(http.StatusOK, posts)
	}
}
