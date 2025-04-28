package middleware

import (
	"net/http"
	"strings"
)

// The suffixes of files that should be blocked from being served by the file server.
// This is case insensitive.
var blacklist = []string{
	"README.md",
	".env",
}

// FileServerBlacklist is a middleware that blocks requests for the suffixes listed in the blacklist to reach the fileserver.

func FileServerBlacklist(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Block access to sensitive files
		path := r.URL.Path
		for _, pattern := range blacklist {
			if strings.HasPrefix(path, "/static/") && hasCaseInsensitiveSuffix(path, pattern) {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}

func hasCaseInsensitiveSuffix(s, suffix string) bool {
	return strings.HasSuffix(strings.ToLower(s), strings.ToLower(suffix))
}

