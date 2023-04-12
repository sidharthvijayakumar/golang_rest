package middleware

import (
	// package to encode and decode the json into struct and vice versa

	// used to read the environment variable

	_ "github.com/lib/pq" // postgres golang driver
)

// response format
type Response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}
