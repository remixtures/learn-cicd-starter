package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
        // Define test cases with their expected outcomes
        tests := []struct {
                name           string
                headers        http.Header
                expectedKey    string
                expectedErr    error
        }{
                {
                        name:           "No Authorization Header",
                        headers:        http.Header{}, // No headers at all
                        expectedKey:    "",
                        expectedErr:    ErrNoAuthHeaderIncluded,
                },
                {
                        name:           "Malformed Authorization Header",
                        headers:        http.Header{"Authorization": []string{"ApiKey"}},
                        expectedKey:    "",
                        expectedErr:    errors.New("malformed authorization header"),
                },
                {
                        name:           "Correct Authorization Header",
                        headers:        http.Header{"Authorization": []string{"ApiKey 12345"}},
                        expectedKey:    "12345",
                        expectedErr:    nil,
                },
                {
                        name:           "Wrong Authorization Scheme",
                        headers:        http.Header{"Authorization": []string{"Bearer 12345"}},
                        expectedKey:    "",
                        expectedErr:    errors.New("malformed authorization header"),
                },
        }

        // Iterate over the test cases and run the function for each
        for _, tt := range tests {
                t.Run(tt.name, func(t *testing.T) {
                        apiKey, err := GetAPIKey(tt.headers)

                        // Check if the error matches the expected error
                        if err != nil && err.Error() != tt.expectedErr.Error() {
                                t.Errorf("expected error %v, got %v", tt.expectedErr, err)
                        }

                        // Check if the API key matches the expected key
                        if apiKey != tt.expectedKey {
                                t.Errorf("expected API key %v, got %v", tt.expectedKey, apiKey)
                        }
                })
        }
}