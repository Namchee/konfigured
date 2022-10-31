package constant

import "errors"

// Configuration error
var (
	ErrMissingToken = errors.New("[Configuration] Missing GitHub access token")
)

// Metadata error
var (
	ErrMalformedMetadata = errors.New("[Metadata] Malformed repository metadata")
)
