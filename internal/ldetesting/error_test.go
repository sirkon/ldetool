package ldetesting

import (
	"os/exec"
	"testing"
)

// TestError ensures these rules cannot be translated into Go code
func TestError(t *testing.T) {
	invalidRules := []string{
		"error_1.lde",
	}

	for _, rule := range invalidRules {
		t.Run(rule, func(t *testing.T) {
			cmd := exec.Command("ldetool", rule)
			cmd.Stderr = &testingWriter{t: t}
			if err := cmd.Run(); err != nil {
				return
			}

			t.Errorf("file %s is expected to raise ldetool error", rule)
		})
	}
}

type testingWriter struct {
	t *testing.T
}

func (w *testingWriter) Write(p []byte) (n int, err error) {
	w.t.Log("\r" + string(p))

	return len(p), nil
}
