package request

import (
	"net/http"
)

type RequestContext struct {
	Request *http.Request
}
