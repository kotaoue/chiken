package main

import (
	"bytes"
	"image/color"
	"io"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

// testWithFlags creates an isolated environment for testing with flags
func testWithFlags(t *testing.T, args []string, testFunc func() error) {
	// Create a new command instance for testing
	cmd := &cobra.Command{
		Use: "chiken",
		RunE: func(cmd *cobra.Command, args []string) error {
			if dump {
				return reOutputs()
			}
			return output()
		},
	}

	// Add flags
	cmd.Flags().StringVarP(&theme, "theme", "t", defaultTheme, "theme color of rooster")
	cmd.Flags().StringVarP(&style, "style", "s", defaultStyle, "style of rooster")
	cmd.Flags().StringVarP(&format, "format", "f", defaultFormat, "format of output image")
	cmd.Flags().StringVarP(&effect, "effect", "e", defaultEffect, "set visual effects")
	cmd.Flags().StringVarP(&background, "background", "b", defaultBackground, "background color. set with hex. example #ffffff. empty is transparent")
	cmd.Flags().StringVarP(&name, "name", "n", defaultName, "name of output image")
	cmd.Flags().IntVarP(&multiple, "multiple", "m", defaultMultiple, "value to be multiplied by 32")
	cmd.Flags().IntVarP(&delay, "delay", "d", defaultDelay, "delay time for gif")
	cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "printing verbose output")
	cmd.Flags().BoolVar(&dump, "dump", false, "re encode from Args Example on README")

	// Reset flags to defaults
	theme = defaultTheme
	style = defaultStyle
	format = defaultFormat
	effect = defaultEffect
	background = defaultBackground
	name = defaultName
	multiple = defaultMultiple
	delay = defaultDelay
	verbose = false
	dump = false

	// Parse flags
	cmd.SetArgs(args[1:])
	err := cmd.ParseFlags(args[1:])
	require.NoError(t, err)

	err = testFunc()
	assert.NoError(t, err)
}

func TestMainFunc(t *testing.T) {
	// Backup original stdout
	origStdout := os.Stdout
	defer func() { os.Stdout = origStdout }()
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Test with no arguments
	testWithFlags(t, []string{"chiken"}, func() error {
		return output()
	})

	// Test with -v flag
	testWithFlags(t, []string{"chiken", "-v"}, func() error {
		return output()
	})

	// Test with output
	testWithFlags(t, []string{"chiken", "-n", "test"}, func() error {
		defer os.Remove("img/test.png")
		return output()
	})

	// Test with re-output
	testWithFlags(t, []string{"chiken", "--dump"}, func() error {
		return reOutputs()
	})

	// Capture and restore stdout
	w.Close()
	_, err := io.ReadAll(r)
	require.NoError(t, err)
}

func TestMainFunction(t *testing.T) {
	// Test with valid arguments
	testWithFlags(t, []string{"chiken", "-m", "1", "-s", "basic", "-t", "white", "-e", "mirror", "-b", "#000000", "-f", "png", "-d", "10", "-n", "test"}, func() error {
		defer os.Remove("img/test.png")
		return output()
	})

	// Test with invalid format
	testWithFlags(t, []string{"chiken", "-f", "invalid"}, func() error {
		err := output()
		assert.Error(t, err, "output() with invalid format should have failed")
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
	require.NoError(t, err)
	assert.NotEmpty(t, buf.String(), "output() should have printed something")

	// Test with invalid background color
	testWithFlags(t, []string{"chiken", "-b", "invalid"}, func() error {
		err := output()
		assert.Error(t, err, "output() with invalid background should have failed")
		return nil
	})
}

func TestEncode(t *testing.T) {
	bgColor := &color.RGBA{R: 0, G: 0, B: 0, A: 255}

	// Test PNG encoding
	testWithFlags(t, []string{"chiken", "-f", "png"}, func() error {
		err := encode(bgColor)
		assert.NoError(t, err, "encode() for png should not fail")
		os.Remove("img/*.png")
		return nil
	})

	// Test GIF encoding
	testWithFlags(t, []string{"chiken", "-f", "gif"}, func() error {
		err := encode(bgColor)
		assert.NoError(t, err, "encode() for gif should not fail")
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
	require.NoError(t, err)
	assert.NotEmpty(t, buf.String(), "printReference() should have printed something")
}

func TestPrintArgs(t *testing.T) {
	testWithFlags(t, []string{"chiken", "-t", "black", "-s", "walk"}, func() error {
		result := printArgs()
		assert.NotEmpty(t, result, "printArgs() should return non-empty string when flags are set")
		assert.Contains(t, result, "-t=black", "printArgs() should contain -t=black")
		assert.Contains(t, result, "-s=walk", "printArgs() should contain -s=walk")
		return nil
	})

	// Test with default values
	testWithFlags(t, []string{"chiken"}, func() error {
		result := printArgs()
		// With default values, should return empty string
		assert.Empty(t, result, "printArgs() with defaults should return empty string")
		return nil
	})
}

func TestFileName(t *testing.T) {
	// Test with no output flag
	testWithFlags(t, []string{"chiken"}, func() error {
		filename := fileName()
		assert.NotEmpty(t, filename, "fileName() should not be empty")
		return nil
	})

	// Test with output flag
	testWithFlags(t, []string{"chiken", "-n", "test"}, func() error {
		filename := fileName()
		assert.Contains(t, filename, "test", "fileName() should contain test")
		return nil
	})
}

func TestCheckFormat(t *testing.T) {
	err := checkFormat("png")
	assert.NoError(t, err, "checkFormat('png') should not fail")

	err = checkFormat("gif")
	assert.NoError(t, err, "checkFormat('gif') should not fail")

	err = checkFormat("invalid")
	assert.Error(t, err, "checkFormat('invalid') should fail")
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
	require.NoError(t, err)
	assert.NotEmpty(t, buf.String(), "reOutputs() should have printed something")
}
