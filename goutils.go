package goutils

import (
	"net/http"
	"strings"
)


// GetClientIP return the client IP of a http request, by the order
// of a given list of headers. If no ip is found in headers, then return request's
// RemoteAddr. This is useful when there are proxy servers between the client and the backend server.
// Example:
// GetClientIP(r,"x-real-ip","x-forwarded-for"), will first check header "x-real-ip"
// if it exists, then split it by "," and return the first part. Otherwise, it will check
// the header "x-forwarded-for" if it exists, then split it by "," and return the first part.
// Otherwise it will return request's RemoteAddr.
//
func GetClientIP(r *http.Request, headers ...string) string {
	for _, header := range headers {
		ip := r.Header.Get(header)
		if ip != "" {
			return strings.Split(ip, ",")[0]
		}
	}
	return strings.Split(r.RemoteAddr, ":")[0]
}
