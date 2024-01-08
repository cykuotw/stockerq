package log_test

import (
	"os"
	"stockerq/configs"
	"stockerq/util/log"
	"testing"
)

func TestServerLogger(t *testing.T) {
	type testcase struct {
		name            string
		event           string
		api             string
		method          string
		body            map[string]string
		fileEnabled     bool
		terminalEnabled bool
		expectFail      bool
		expectError     error
	}

	subtests := []testcase{
		{
			name:   "valid file enabled only",
			event:  "test event",
			api:    "/api/testcase1",
			method: "GET",
			body: map[string]string{
				"id": "testid",
			},
			fileEnabled:     true,
			terminalEnabled: false,
			expectFail:      false,
			expectError:     nil,
		},
		{
			name:   "valid terminal enabled only",
			event:  "test event",
			api:    "/api/testcase1",
			method: "GET",
			body: map[string]string{
				"id": "testid",
			},
			fileEnabled:     false,
			terminalEnabled: true,
			expectFail:      false,
			expectError:     nil,
		},
	}

	for _, test := range subtests {
		t.Run(test.name, func(t *testing.T) {
			configs.Init()
			en := ""
			if test.fileEnabled {
				en = "True"
			} else {
				en = "False"
			}
			os.Setenv("logServerTermEn", en)
			if test.terminalEnabled {
				en = "True"
			} else {
				en = "False"
			}
			os.Setenv("logServerFileEn", en)

			log.LogServer(test.event, test.api, test.method, test.body)
		})
	}
}
