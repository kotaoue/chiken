package main

import (
	"bytes"
	"flag"
	"image/color"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

// testWithFlags creates an isolated environment for testing with flags
func testWithFlags(t *testing.T, args []string, testFunc func() error) {
	// Backup original os.Args and CommandLine
	origArgs := os.Args
	origCommandLine := flag.CommandLine
	defer func() {
		os.Args = origArgs
		flag.CommandLine = origCommandLine
	}()

	// Set up new environment
	os.Args = args
	flag.CommandLine = flag.NewFlagSet(args[0], flag.ContinueOnError)
	setupFlags()
	if err := flag.CommandLine.Parse(args[1:]); err != nil {
		t.Fatalf("Failed to parse flags: %v", err)
	}

	if err := testFunc(); err != nil {
		t.Errorf("Test function failed: %v", err)
	}
}

func TestMainFunc(t *testing.T) {
	// Backup original stdout
	origStdout := os.Stdout
	defer func() { os.Stdout = origStdout }()
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Test with no arguments
	testWithFlags(t, []string{"chiken"}, func() error {
		return Main()
	})

	// Test with -v flag
	testWithFlags(t, []string{"chiken", "-v"}, func() error {
		return Main()
	})

	// Test with output
	testWithFlags(t, []string{"chiken", "-n", "test"}, func() error {
		defer os.Remove("img/test.png")
		return Main()
	})

	// Test with re-output
	testWithFlags(t, []string{"chiken", "-dump"}, func() error {
		return Main()
	})

	// Capture and restore stdout
	w.Close()
	_, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMainFunction(t *testing.T) {
	// Test with valid arguments
	testWithFlags(t, []string{"chiken", "-m", "1", "-s", "basic", "-t", "white", "-e", "mirror", "-b", "#000000", "-f", "png", "-d", "10", "-n", "test"}, func() error {
		defer os.Remove("img/test.png")
		return Main()
	})

	// Test with invalid format
	testWithFlags(t, []string{"chiken", "-f", "invalid"}, func() error {
		if err := Main(); err == nil {
			t.Error("Main() with invalid format should have failed")
		}
		return nil
	})
}

func TestOutput(t *testing.T) {
	// Backup original stdout
	origStdout := os.Stdout
	defer func() { os.Stdout = origStdout }()
	r, w, _ := os.Pipe()
	os.Stdout = w

	testWithFlags(t, []string{"chiken"}, func() error {
		output()
		return nil
	})

	w.Close()
	var buf bytes.Buffer
	_, err := buf.ReadFrom(r)
	if err != nil {
		t.Fatal(err)
	}
	if buf.String() == "" {
		t.Error("output() should have printed something")
	}

	// Test with invalid background color
	testWithFlags(t, []string{"chiken", "-b", "invalid"}, func() error {
		if err := output(); err == nil {
			t.Error("output() with invalid background should have failed")
		}
		return nil
	})
}

func TestEncode(t *testing.T) {
	bgColor := &color.RGBA{R: 0, G: 0, B: 0, A: 255}

	// Test PNG encoding
	testWithFlags(t, []string{"chiken", "-f", "png"}, func() error {
		if err := encode(bgColor); err != nil {
			t.Errorf("encode() for png failed: %v", err)
		}
		os.Remove("img/*.png")
		return nil
	})

	// Test GIF encoding
	testWithFlags(t, []string{"chiken", "-f", "gif"}, func() error {
		if err := encode(bgColor); err != nil {
			t.Errorf("encode() for gif failed: %v", err)
		}
		os.Remove("img/*.gif")
		return nil
	})
}

func TestPrintReference(t *testing.T) {
	// Backup original stdout
	origStdout := os.Stdout
	defer func() { os.Stdout = origStdout }()
	r, w, _ := os.Pipe()
	os.Stdout = w

	printReference()

	w.Close()
	var buf bytes.Buffer
	_, err := buf.ReadFrom(r)
	if err != nil {
		t.Fatal(err)
	}
	if buf.String() == "" {
		t.Error("printReference() should have printed something")
	}
}

func TestPrintArgs(t *testing.T) {
	testWithFlags(t, []string{"chiken", "-t", "black", "-s", "walk"}, func() error {
		result := printArgs()
		if result == "" {
			t.Error("printArgs() should return non-empty string when flags are set")
		}
		if !strings.Contains(result, "-t=black") || !strings.Contains(result, "-s=walk") {
			t.Errorf("printArgs() = %s, want to contain -t=black and -s=walk", result)
		}
		return nil
	})

	// Test with default values
	testWithFlags(t, []string{"chiken"}, func() error {
		result := printArgs()
		// With default values, should return empty string
		if result != "" {
			t.Errorf("printArgs() with defaults = %s, want empty string", result)
		}
		return nil
	})
}

func TestFileName(t *testing.T) {
	// Test with no output flag
	testWithFlags(t, []string{"chiken"}, func() error {
		filename := fileName()
		if filename == "" {
			t.Error("fileName() should not be empty")
		}
		return nil
	})

	// Test with output flag
	testWithFlags(t, []string{"chiken", "-n", "test"}, func() error {
		filename := fileName()
		if !strings.Contains(filename, "test") {
			t.Errorf("fileName() = %s, want to contain test", filename)
		}
		return nil
	})
}

func TestCheckFormat(t *testing.T) {
	if err := checkFormat("png"); err != nil {
		t.Errorf("checkFormat('png') failed: %v", err)
	}
	if err := checkFormat("gif"); err != nil {
		t.Errorf("checkFormat('gif') failed: %v", err)
	}
	if err := checkFormat("invalid"); err == nil {
		t.Error("checkFormat('invalid') should have failed")
	}
}

func TestReOutputs(t *testing.T) {
	// Backup original stdout
	origStdout := os.Stdout
	defer func() { os.Stdout = origStdout }()
	r, w, _ := os.Pipe()
	os.Stdout = w

	reOutputs()

	w.Close()
	var buf bytes.Buffer
	_, err := buf.ReadFrom(r)
	if err != nil {
		t.Fatal(err)
	}
	if buf.String() == "" {
		t.Error("reOutputs() should have printed something")
	}
}
