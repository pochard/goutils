# goutils
This is a golang utils module with some common functions.

Import it in your program as:
```go
      import "github.com/pochard/goutils"
```

## API
### func GetClientIP(r *http.Request, headers ...string) string

## Example
``` go
// GetClientIP return the client IP of a http request, by the order
// of a given list of headers. If no ip is found in headers, then return request's
// RemoteAddr. This is useful when there are proxy servers between the client and the backend server.
// Example:
// GetClientIP(r,"x-real-ip","x-forwarded-for"), will first check header "x-real-ip"
// if it exists, then split it by "," and return the first part. Otherwise, it will check
// the header "x-forwarded-for" if it exists, then split it by "," and return the first part.
// Otherwise it will return request's RemoteAddr.
//
func handler(w http.ResponseWriter, r *http.Request) {
	clientIP := goutils.GetClientIP(r,"x-real-ip","x-forwarded-for")
}

```