package commons

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/kitabisa/golang-poc/config"
	phttp "github.com/kitabisa/perkakas/v2/http"
	"github.com/kitabisa/perkakas/v2/structs"
)

var cfg = config.Config()

// InjectErrors injecting all error response to the handler context
func InjectErrors(handlerCtx *phttp.HttpHandlerContext) {
	handlerCtx.AddError(ErrDBConn, ErrDBConnResp)
	handlerCtx.AddError(ErrCacheConn, ErrCacheConnResp)
	// etc...
}

// getErrorResponce will return error response code & description object according to error code
func getErrorResponce(errorCode string) structs.Response {
	return structs.Response{
		ResponseCode: errorCode,
		ResponseDesc: structs.ResponseDesc{
			ID: cfg.GetString(fmt.Sprintf("RESPONSE_CODE_ID_%s", errorCode)),
			EN: cfg.GetString(fmt.Sprintf("RESPONSE_CODE_EN_%s", errorCode)),
		},
	}
}

// ErrDBConn error type for Error DB Connection
var ErrDBConn = errors.New("ErrDBConn")

// ErrDBConnResp ErrDBConn's response
var ErrDBConnResp *structs.ErrorResponse = &structs.ErrorResponse{
	Response:   getErrorResponce("101001"),
	HttpStatus: http.StatusInternalServerError,
}

// ErrCacheConn error type for Error Cache Connection
var ErrCacheConn = errors.New("ErrCacheConn")

// ErrCacheConnResp ErrCacheConn's response
var ErrCacheConnResp *structs.ErrorResponse = &structs.ErrorResponse{
	Response:   getErrorResponce("101002"),
	HttpStatus: http.StatusInternalServerError,
}
