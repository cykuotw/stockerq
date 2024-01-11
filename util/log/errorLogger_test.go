package log_test

import (
	"fmt"
	"math/rand"
	"os"
	"stockerq/configs"
	"stockerq/util/log"
	apperror "stockerq/web/app/app-error"
	"testing"
)

func TestErrorLogger(t *testing.T) {
	type testcase struct {
		name            string
		layer           string
		apperror        apperror.Error
		fileEnabled     bool
		terminalEnabled bool
		expectFail      bool
		expectError     error
	}

	subtests := []testcase{
		{
			name:            "valid file enabled only",
			fileEnabled:     true,
			terminalEnabled: false,
			expectFail:      false,
			expectError:     nil,
		},
		{
			name:            "valid terminal enabled only",
			fileEnabled:     false,
			terminalEnabled: true,
			expectFail:      false,
			expectError:     nil,
		},
	}

	var getRandomErr = func() (string, apperror.Error) {
		randNum := rand.Uint32() % 3
		err := fmt.Errorf("test error")
		switch randNum {
		case 0:
			errC := apperror.NewControllerError(err)
			return errC.Layer, errC

		case 1:
			errM := apperror.NewModelError(err)
			return errM.Layer, errM

		case 2:
			errR := apperror.NewRoutingError(err)
			return errR.Layer, errR
		}

		return "", nil
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
			os.Setenv("logErrorFileEn", en)
			if test.terminalEnabled {
				en = "True"
			} else {
				en = "False"
			}
			os.Setenv("logErrorTermEn", en)

			layer, err := getRandomErr()
			log.LogError(layer, err)
		})
	}
}
