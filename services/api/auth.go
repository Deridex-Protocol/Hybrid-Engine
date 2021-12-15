package api

import (
	"net/http"
	"strings"
	"time"

	"bitbucket.ideasoft.io/dex/dex-backend/common"
	"bitbucket.ideasoft.io/dex/dex-backend/common/crypto"
	"github.com/labstack/echo/v4"
)

const addressContextKey = "Authentication-Address"

func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		ctx.Response().Header().Set(echo.HeaderServer, "Echo/3.0")
		hydroAuthToken := ctx.Request().Header.Get(common.AuthenticationHeaderKey)

		hydroAuthTokens := strings.Split(hydroAuthToken, "#")
		if len(hydroAuthTokens) != 3 {
			return echo.NewHTTPError(http.StatusUnauthorized, "Authentication should be like {address}#Authentication:2006-01-02T15:04:05Z07:00#{signature}")
		}

		if !strings.Contains(hydroAuthTokens[1], "Authentication") {
			return echo.NewHTTPError(http.StatusUnauthorized, "Authentication message should be like Authentication:2006-01-02T15:04:05Z07:00")
		}

		messageTime, err := time.Parse(time.RFC3339, strings.Trim(hydroAuthTokens[1], "Authentication:"))
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Authentication message should be like Authentication:2006-01-02T15:04:05Z07:00")
		}

		if time.Now().After(messageTime) {
			return echo.NewHTTPError(http.StatusUnauthorized, "Authentication token expired")
		}

		valid, err := crypto.IsValidAuthSignature(hydroAuthTokens[0], hydroAuthTokens[1], hydroAuthTokens[2])
		if err != nil || !valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Authentication valid failed, please check your authentication")
		}

		ctx.Set(addressContextKey, strings.ToLower(hydroAuthTokens[0]))
		return next(ctx)
	}
}
