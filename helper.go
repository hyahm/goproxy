package goproxy

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"strings"
)

var inOurTests bool // whether we're in our own tests
func shouldPanicOnCopyError(req *http.Request) bool {
	if inOurTests {
		// Our tests know to handle this panic.
		return true
	}
	if req.Context().Value(http.ServerContextKey) != nil {
		// We seem to be running under an HTTP server, so
		// it'll recover the panic.
		return true
	}
	// Otherwise act like Go 1.10 and earlier to not break
	// existing tests.
	return false
}



func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}


func Md5 (data []byte) string {
	m := md5.New()
	m.Write(data)
	return fmt.Sprintf("%x", m.Sum(nil))
}