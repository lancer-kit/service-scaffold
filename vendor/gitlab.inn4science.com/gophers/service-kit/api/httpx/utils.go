package httpx

import (
	"fmt"
	"net/http"
	"sort"
	"strings"
)

// messageForSigning concatenates passed request data in a fixed format.
func messageForSigning(service, method, url, body, authHeaders string) string {
	return fmt.Sprintf("service:%s;method:%s;path:%s;authHeaders:%s;body:%s;",
		service, method, url, authHeaders, body)
}

// addCookies sets all `cookies` to the http.Request.
func addCookies(r *http.Request, cookies []*http.Cookie) *http.Request {
	for _, k := range cookies {
		r.AddCookie(k)
	}
	return r
}

//headersForSigning concatenates passed keys from headers map
func headersForSigning(headers map[string]string) string {
	var keys []string
	for key := range headers {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return strings.Join(keys, ":")
}
