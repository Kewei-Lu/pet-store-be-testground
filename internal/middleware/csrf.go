package middleware

import (
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/grand"
	"net/http"
	"strings"
)

var (
	TokenCookieName string = "X-Token"
	TokenRequestKey string = "X-Token"
	TokenLength     int    = 32
)

// NewWithCfg creates and returns a CSRF middleware with incoming configuration.
func NewCsrf() func(r *ghttp.Request) {
	return func(r *ghttp.Request) {

		// Read the token in the request cookie
		tokenInCookie := r.Cookie.Get(TokenCookieName).String()
		if tokenInCookie == "" {
			// Generate a random token
			tokenInCookie = grand.S(TokenLength)
		}

		// Read the token attached to the request
		// Read priority: Router < Query < Body < Form < Custom < Header
		tokenInRequestData := r.Header.Get(TokenRequestKey)
		if tokenInRequestData == "" {
			tokenInRequestData = r.GetRequest(TokenRequestKey).String()
		}

		switch r.Method {
		case http.MethodGet, http.MethodHead, http.MethodOptions, http.MethodTrace:
			// No verification required
		default:
			// Authentication token
			if !strings.EqualFold(tokenInCookie, tokenInRequestData) {
				fmt.Println("tokenInCookie: ", tokenInCookie, "tokenInRequestData: ", tokenInRequestData)
				r.Response.WriteStatusExit(http.StatusForbidden, "CSRF Rejection")
				return
			}
		}
		r.Middleware.Next()
	}
}
