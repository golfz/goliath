package contexts

import (
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"github.com/golfz/goliath/cleanarch/data/output"
)

type authContext struct {
	request *http.Request
}

type AuthContextor interface {
	GetAuthorizationHeader() (string, output.GoliathError)
	GetBearerToken() (string, output.GoliathError)
}

func (ctx *authContext) GetAuthorizationHeader() (string, output.GoliathError) {
	auth := ctx.request.Header.Get("Authorization")

	if strings.TrimSpace(auth) == "" {
		return "", &output.Error{
			Status:  http.StatusUnauthorized,
			Time:    time.Now(),
			Type:    "goliath",
			Code:    "goliath.context.GetAuthorizationHeader.no-auth-header",
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

func (ctx *authContext) GetBearerToken() (string, output.GoliathError) {
	rawToken, err := ctx.GetAuthorizationHeader()
	if err != nil {
		return "", err
	}

	rawToken = strings.TrimSpace(rawToken)

	if !isBearerAuth(rawToken) {
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

	rawToken = strings.Replace(rawToken, bearerStartPattern, "", 1)

	return rawToken, nil
}
