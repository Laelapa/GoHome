package logging

const (

	// HTTP request fields ----------------------------------------------------

	// FieldRemoteAddr is the log field name for the remote address of an HTTP request.
	FieldRemoteAddr = "remote_addr"

	// FieldMethod is the log field name for the HTTP method used in a request.
	FieldMethod = "method"

	// FieldPath is the log field name for the requested URL path.
	FieldPath = "path"

	// FieldReferer is the log field name for the HTTP referer header.
	FieldReferer = "referer"

	// FieldUserAgent is the log field name for the HTTP user agent header.
	FieldUserAgent = "user_agent"

	// HTTP response fields ---------------------------------------------------

	// FieldStatusCode is the log field name for the HTTP response status code.
	FieldStatusCode = "status_code"

	// FieldResponseTime is the log field name for the time taken to respond to a request.
	FieldResponseTime = "response_time"

	// Error fields -----------------------------------------------------------

	// FieldError is the log field name for error information.
	FieldError = "error"

	// FieldService is the log field name for the service identifier.
	FieldService = "service"

	// FieldEnvironment is the log field name for the application environment.
	FieldEnvironment = "environment"

	// FieldVersion is the log field name for the application version.
	FieldVersion = "version"

	// FieldServerAddr is the log field name for the server's address.
	FieldServerAddr = "server_addr"

	// Logging fields ---------------------------------------------------------

	// FieldLoggingLevel is the log field name for the logging level.
	FieldLoggingLevel = "logging_level"

	// Other fields -----------------------------------------------------------

	// FieldDuration is the log field name for durations in milliseconds.
	FieldDuration = "duration_ms"
)
