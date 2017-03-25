package middleware

import "github.com/jacksontoomey/gofrit/request"

// Middleware interface provides one function to pre and post process requests
type Middleware interface {
	BeforeRequest(context *request.RequestContext)
	AfterRequest(context *request.RequestContext)
}
