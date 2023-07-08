package middleware

import (
	"github.com/serve-it-yourself/lux/context"
	"net/http"
)

type AuthChecker func(lc *context.LuxContext, authorizationHeader string, tokenCookies ...*http.Cookie) bool

func Auth(authChecker AuthChecker, tokenName ...string) Set {
	return Set{
		Request: func(ctx *context.LuxContext) (*context.LuxContext, int) {
			req := ctx.Request
			authorizationHeader := req.Header.Get("Authorization")
			cookies := []*http.Cookie(nil)
			for _, name := range tokenName {
				cookie, err := req.Cookie(name)
				if err == nil {
					cookies = append(cookies, cookie)
				}
			}
			if authChecker(ctx, authorizationHeader, cookies...) {
				return ctx, http.StatusOK
			}
			return ctx, http.StatusUnauthorized
		},
		Response: nil,
	}
}
