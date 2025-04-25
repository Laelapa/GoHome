package logging

const (

	// HTTP request fields
	FieldRemoteAddr = "remote_addr"
	FieldMethod     = "method"
	FieldPath       = "path"
	FieldReferer    = "referer"
	FieldUserAgent  = "user_agent"

	// HTTP response fields
	FieldStatusCode   = "status_code"
	FieldResponseTime = "response_time"

	// Error fields
	FieldError = "error"

	// Application fields
	FieldService     = "service"
	FieldEnvironment = "environment"
	FieldVersion     = "version"
	FieldServerAddr  = "server_addr"

	// Logging fields
	FieldLoggingLevel = "logging_level"

	// Other fields
	FieldDuration = "duration_ms"
)
