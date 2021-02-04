package contexts

import (
	"github.com/golfz/goliath/cleanarch/data/output"
	"net/http"
	"runtime/debug"
	"strings"
	"time"
)

type authContext struct {
	request *http.Request
}

type AuthContextor interface {
	GetAuthToken() (string, output.GoliathError)
	GetBearerToken(auth string) (string, output.GoliathError)
}

func (ctx *authContext) GetAuthToken() (string, output.GoliathError) {
	auth := ctx.request.Header.Get("authorization")

	if strings.TrimSpace(auth) == "" {
		return "", &output.Error{
			Status:  http.StatusUnauthorized,
			Time:    time.Now(),
			Type:    "goliath",
			Code:    "goliath.context.GetAuthToken.no-auth-header",
			Error:   "empty authorization header",
			Message: "empty authorization header",
			ErrorDev: output.ErrorDev{
				Error:      "empty authorization header",
				Stacktrace: string(debug.Stack()),
			},
		}
	}

	auth = strings.TrimSpace(auth)

	return auth, nil
}

func (ctx *authContext) GetBearerToken(s string) (string, output.GoliathError) {
	s = strings.TrimSpace(s)

	if !isBearerAuth(s) {
		return "", &output.Error{
			Status:  http.StatusUnauthorized,
			Time:    time.Now(),
			Type:    "goliath",
			Code:    "goliath.context.GetBearerToken.no-bearer",
			Error:   "authorization is not bearer",
			Message: "authorization is not bearer",
			ErrorDev: output.ErrorDev{
				Error:      "authorization is not bearer",
				Stacktrace: string(debug.Stack()),
			},
		}
	}

	s = strings.Replace(s, bearerStartPattern, "", 1)

	return s, nil
}
