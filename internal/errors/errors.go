package errors

import (
	"fmt"
	"net/http"

	"github.com/gookit/goutil/arrutil"
)

// 验证错误
var (
	BROWSER_ID_REQUIRED = APIError{Code: http.StatusBadRequest, Message: "id.not.empty"}
)

// 系统错误
var (
	RATE_LIMIT_ERROR = APIError{Code: http.StatusTooManyRequests, Message: "too.many.request"}
	// 系统未知异常
	SERVER_ERROR = APIError{Code: http.StatusInternalServerError, Message: "server.unknown.exception"}
)

// 鉴权
var (

	// token过期
	AUTH_TOKEN_UNAUTHORIZED_ACCESS = APIError{Code: http.StatusUnauthorized, Message: "auth.token.expired"}
	//无访问权限
	AUTH_NOT_ACCESS_PERMISSION = APIError{Code: http.StatusUnauthorized, Message: "no.access.permission"}
)

type APIError struct {
	Code    int
	Message string
	Extra   []string
}

func (apiError APIError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", apiError.Code, fmt.Sprintf(apiError.Message, arrutil.StringsToAnys(apiError.Extra)...))
}

func NewAPIError(code int, message string) APIError {
	return APIError{
		Code:    code,
		Message: message,
	}
}
