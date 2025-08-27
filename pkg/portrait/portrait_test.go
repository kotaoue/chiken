package portrait

import (
	"image/color"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPortrait(t *testing.T) {
	opts := Options{
		Size:            32,
		BaseSize:        32,
		Multiple:        1,
		Style:           "basic",
		Theme:           "white",
		Effect:          "mirror",
		BackgroundColor: &color.RGBA{R: 0, G: 0, B: 0, A: 255},
		Format:          "png",
		Delay:           10,
		FileName:        "test.png",
		Verbose:         true,
	}
	p := NewPortrait(opts)
	assert.NotNil(t, p, "NewPortrait() should not return nil")
	assert.Equal(t, opts, p.opt, "NewPortrait() should set options correctly")
}

func TestPortrait_Encode(t *testing.T) {
	// Test PNG encoding
	optsPng := Options{
		Size:            32,
		BaseSize:        32,
		Multiple:        1,
		Style:           "basic",
		Theme:           "white",
		Effect:          "mirror",
		BackgroundColor: &color.RGBA{R: 0, G: 0, B: 0, A: 255},
		Format:          "png",
		FileName:        "test.png",
	}
	pPng := NewPortrait(optsPng)
	err := pPng.Encode()
	assert.NoError(t, err, "Portrait.Encode() for png should not fail")
	os.Remove("test.png")

	// Test GIF encoding
	optsGif := Options{
		Size:            32,
		BaseSize:        32,
		Multiple:        1,
		Style:           "basic",
		Theme:           "white",
		Effect:          "mirror",
		BackgroundColor: &color.RGBA{R: 0, G: 0, B: 0, A: 255},
		Format:          "gif",
		FileName:        "test.gif",
	}
	pGif := NewPortrait(optsGif)
	err = pGif.Encode()
	assert.NoError(t, err, "Portrait.Encode() for gif should not fail")
	os.Remove("test.gif")
}

func TestVPrint(t *testing.T) {
	verbose = true
	// Since this function just wraps fmt.Print, we can't easily test its output.
	// We'll just call it to ensure it doesn't panic.
	assert.NotPanics(t, func() {
		vPrint("test")
	}, "vPrint should not panic")
}

func TestVPrintln(t *testing.T) {
	verbose = true
	// Since this function just wraps fmt.Println, we can't easily test its output.
	// We'll just call it to ensure it doesn't panic.
	assert.NotPanics(t, func() {
		vPrintln("test")
	}, "vPrintln should not panic")
}

func TestVPrintf(t *testing.T) {
	verbose = true
	// Since this function just wraps fmt.Printf, we can't easily test its output.
	// We'll just call it to ensure it doesn't panic.
	assert.NotPanics(t, func() {
		vPrintf("test %s", "value")
	}, "vPrintf should not panic")
}

func TestPortrait_Encode_Error(t *testing.T) {
	// Test invalid file path for PNG
	optsPng := Options{
		Size:            32,
		BaseSize:        32,
		Multiple:        1,
		Style:           "basic",
		Theme:           "white",
		BackgroundColor: &color.RGBA{R: 0, G: 0, B: 0, A: 255},
		Format:          "png",
		FileName:        "invalid/path/test.png",
	}
	pPng := NewPortrait(optsPng)
	err := pPng.Encode()
	assert.Error(t, err, "Portrait.Encode() for png with invalid path should fail")

	// Test invalid file path for GIF
	optsGif := Options{
		Size:            32,
		BaseSize:        32,
		Multiple:        1,
		Style:           "basic",
		Theme:           "white",
		BackgroundColor: &color.RGBA{R: 0, G: 0, B: 0, A: 255},
		Format:          "gif",
		FileName:        "invalid/path/test.gif",
	}
	pGif := NewPortrait(optsGif)
	err = pGif.Encode()
	assert.Error(t, err, "Portrait.Encode() for gif with invalid path should fail")

	// Test invalid style
	optsInvalidStyle := Options{
		Style: "invalid",
	}
	pInvalidStyle := NewPortrait(optsInvalidStyle)
	err = pInvalidStyle.encodePng()
	assert.Error(t, err, "encodePng with invalid style should fail")

	// Test invalid theme
	optsInvalidTheme := Options{
		Style: "basic",
		Theme: "invalid",
	}
	pInvalidTheme := NewPortrait(optsInvalidTheme)
	err = pInvalidTheme.encodePng()
	assert.Error(t, err, "encodePng with invalid theme should fail")
}
