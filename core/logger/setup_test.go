package logger

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestGetLoggerDefault(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to open pipe: %v", err)
	}

	stderr := os.Stderr
	os.Stderr = w
	defer func() {
		os.Stderr = stderr
	}()

	// Call SetupLogging again so it picks up stderr change
	SetupLogging(Config{Stderr: true})
	log := getLogger("test")

	log.Error("scooby")
	w.Close()

	buf := &bytes.Buffer{}
	if _, err := io.Copy(buf, r); err != nil && err != io.ErrClosedPipe {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.Contains(buf.String(), "scooby") {
		t.Errorf("got %q, wanted it to contain log output", buf.String())
	}
}

func TestCustomCore(t *testing.T) {
	r1, w1, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to open pipe: %v", err)
	}
	r2, w2, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to open pipe: %v", err)
	}

	// logging should work with the custom core
	SetPrimaryCore(newCore(PlaintextOutput, w1, LevelDebug))
	log := getLogger("test")
	log.Error("scooby")

	// SetPrimaryCore should replace the core in previously created loggers
	SetPrimaryCore(newCore(PlaintextOutput, w2, LevelDebug))
	log.Error("doo")

	w1.Close()
	w2.Close()

	buf1 := &bytes.Buffer{}
	buf2 := &bytes.Buffer{}
	if _, err := io.Copy(buf1, r1); err != nil && err != io.ErrClosedPipe {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, err := io.Copy(buf2, r2); err != nil && err != io.ErrClosedPipe {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(buf1.String(), "scooby") {
		t.Errorf("got %q, wanted it to contain log output", buf1.String())
	}
	if !strings.Contains(buf2.String(), "doo") {
		t.Errorf("got %q, wanted it to contain log output", buf2.String())
	}
}
