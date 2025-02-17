package myMiddleware

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/x/errors"
	"net/http"
)

func PanicMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				httpx.Error(w, errors.New(50000, r.(string)))
			}
		}()
		next(w, r)
	}
}
