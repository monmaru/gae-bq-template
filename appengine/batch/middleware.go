package batch

import (
	"context"
	"fmt"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func withAppHandler(h appHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)

		if ae := h(ctx, w, r); ae != nil {
			msg := fmt.Sprintf("Handler error: status code: %d, message: %s, underlying err: %#v",
				ae.Code, ae.Message, ae.Error)
			http.Error(w, msg, ae.Code)
			log.Errorf(ctx, "%s", msg)
		}
	}
}

type appHandler func(context.Context, http.ResponseWriter, *http.Request) *appError

type appError struct {
	Error   error
	Message string
	Code    int
}

func appErrorf(err error, statusCode int, format string, v ...interface{}) *appError {
	return &appError{
		Error:   err,
		Message: fmt.Sprintf(format, v...),
		Code:    statusCode,
	}
}
