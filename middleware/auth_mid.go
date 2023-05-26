package middleware

import (
	"net/http"
	"vick1208/restful-api/helper"
	"vick1208/restful-api/model/web"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middle *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if "RAHASIA" == r.Header.Get("X-API-Key") {
		// ok
		middle.Handler.ServeHTTP(w, r)
	} else {
		// err
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResp := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "UNAUTHORIZED",
		}
		helper.WriteToRespBody(w, webResp)
	}
}
