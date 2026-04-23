package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		inputHeader   string
		expectedValue string
	}{
		{
			name:          "Valid Header",
			inputHeader:   "ApiKey 12345",
			expectedValue: "12345",
		},
		{
			name:          "Empty Header",
			inputHeader:   "",
			expectedValue: "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			headers := http.Header{}
			if tc.inputHeader != "" {
				headers.Add("Authorization", tc.inputHeader)
			}

			got, _ := GetAPIKey(headers)
			if got != tc.expectedValue {
				t.Errorf("failed %s: expected %v, got %v", tc.name, tc.expectedValue, got)
			}
		})
	}
}
