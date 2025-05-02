package logging

import (
	"net/http"
	"strings"
)

const maxHeaderLength int = 1000 // TODO: Potentially make this configurable

// TODO: Consider generalizing by using X-forwarded-for instead of Fly.io specific header.
//
// getClientIP retrieves the client IP address from the request if a reverse proxy is sitting in the middle.
// Currently only works if deployed on fly.io.
// FIXME: SECURITY: Fly-Client-IP can obviously be spoofed in non-fly.io environments and should be sanitized appropriately.
func getClientIP(r *http.Request) string {
	// Check if the request has a fly.io forwarded header
	if clientIP := r.Header.Get("Fly-Client-IP"); clientIP != "" {
		return clientIP
	}
	// Fallback to the remote address
	return "fly.io r-proxy: " + r.RemoteAddr
}

func sanitizeLogValue(v string) string {
	replacer := strings.NewReplacer(

		// Control characters
		"\u0000", "[NUL]", // Null byte - can cause string truncation
		"\u001b", "[ESC]", // ANSI escape sequences

		// Invisible space characters
		"\u200B", "[ZWS]", // Zero width space
		"\u2028", "[LS]", // Line separator
		"\u2029", "[PS]", // Paragraph separator
		"\u2063", "[ICS]", // Invisible separator

		// JSON structural characters handled by zap

		// Newline characters
		"\n", "[LF]", // Line feed
		"\r", "[CR]", // Carriage return
	)
	return replacer.Replace(v)
}

// Protects against flooding the logs with long strings.
//
// maxLength: the maximum length in characters (runes) of the string to log
func truncateLogValue(v string, maxLength int) string {
	vRuned := []rune(v)
	if len(vRuned) <= maxLength {
		return v
	}

	return string(vRuned[:maxLength]) + "... [truncated]"
}

// filetLogValue sanitizes and truncates a string for logging purposes.
func filetLogValue(v string) string {

	return truncateLogValue(sanitizeLogValue(v), maxHeaderLength)

}
