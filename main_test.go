package main

import (
	"bytes"
	"flag"
	"image/color"
	"io/ioutil"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}

func TestMainFunc(t *testing.T) {
	// Backup original os.Args
	origArgs := os.Args
	defer func() { os.Args = origArgs }()

	// Backup original stdout
	origStdout := os.Stdout
	defer func() { os.Stdout = origStdout }()
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Test with no arguments
	os.Args = []string{"chiken"}
	main()

	// Test with -h flag
	os.Args = []string{"chiken", "-h"}
	main()

	// Test with -v flag
	os.Args = []string{"chiken", "-v"}
	main()

	// Test with invalid flag
	os.Args = []string{"chiken", "-invalid"}
	main()

	// Test with reference
	os.Args = []string{"chiken", "-r"}
	main()

	// Test with args
	os.Args = []string{"chiken", "-a"}
	main()

	// Test with output
	os.Args = []string{"chiken", "-o", "test.png"}
	main()
	// Clean up the created file
	os.Remove("test.png")

	// Test with re-output
	os.Args = []string{"chiken", "-re"}
	main()

	// Capture and restore stdout
	w.Close()
	_, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMainFunction(t *testing.T) {
	// Backup original os.Args
	origArgs := os.Args
	defer func() { os.Args = origArgs }()

	// Test with valid arguments
	os.Args = []string{"chiken", "-s", "16", "-bs", "16", "-m", "1", "-st", "basic", "-th", "white", "-e", "mirror", "-bg", "#000000", "-f", "png", "-d", "10", "-o", "test.png"}
	if err := Main(); err != nil {
		t.Errorf("Main() with valid args failed: %v", err)
	}
	os.Remove("test.png")

	// Test with invalid format
	os.Args = []string{"chiken", "-f", "invalid"}
	if err := Main(); err == nil {
		t.Error("Main() with invalid format should have failed")
	}
}

func TestOutput(t *testing.T) {
	// Backup original stdout
	origStdout := os.Stdout
	defer func() { os.Stdout = origStdout }()
	r, w, _ := os.Pipe()
	os.Stdout = w

	output()

	w.Close()
	var buf bytes.Buffer
	_, err := buf.ReadFrom(r)
	if err != nil {
		t.Fatal(err)
	}
	if buf.String() == "" {
		t.Error("output() should have printed something")
	}
}

func TestEncode(t *testing.T) {
	bgColor := &color.RGBA{R: 0, G: 0, B: 0, A: 255}
	// Test PNG encoding
	flag.CommandLine.Set("f", "png")
	if err := encode(bgColor); err != nil {
		t.Errorf("encode() for png failed: %v", err)
	}
	os.Remove("*.png")

	// Test GIF encoding
	flag.CommandLine.Set("f", "gif")
	if err := encode(bgColor); err != nil {
		t.Errorf("encode() for gif failed: %v", err)
	}
	os.Remove("*.gif")
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
	// Backup original stdout
	origStdout := os.Stdout
	defer func() { os.Stdout = origStdout }()
	r, w, _ := os.Pipe()
	os.Stdout = w

	printArgs()

	w.Close()
	var buf bytes.Buffer
	_, err := buf.ReadFrom(r)
	if err != nil {
		t.Fatal(err)
	}
	if buf.String() == "" {
		t.Error("printArgs() should have printed something")
	}
}

func TestFileName(t *testing.T) {
	// Test with no output flag
	flag.CommandLine.Set("o", "")
	name := fileName()
	if name == "" {
		t.Error("fileName() should not be empty")
	}

	// Test with output flag
	flag.CommandLine.Set("o", "test.png")
	name = fileName()
	if name != "test.png" {
		t.Errorf("fileName() = %s, want test.png", name)
	}
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
