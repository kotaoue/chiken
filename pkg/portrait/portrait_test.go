package portrait

import (
	"image/color"
	"os"
	"reflect"
	"testing"
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
	if p == nil {
		t.Fatal("NewPortrait() returned nil")
	}
	if !reflect.DeepEqual(p.opt, opts) {
		t.Errorf("NewPortrait() opts = %v, want %v", p.opt, opts)
	}
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
	if err := pPng.Encode(); err != nil {
		t.Errorf("Portrait.Encode() for png failed: %v", err)
	}
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
	if err := pGif.Encode(); err != nil {
		t.Errorf("Portrait.Encode() for gif failed: %v", err)
	}
	os.Remove("test.gif")
}

func TestVPrint(t *testing.T) {
	verbose = true
	// Since this function just wraps fmt.Print, we can't easily test its output.
	// We'll just call it to ensure it doesn't panic.
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("vPrint panicked: %v", r)
		}
	}()
	vPrint("test")
}

func TestVPrintln(t *testing.T) {
	verbose = true
	// Since this function just wraps fmt.Println, we can't easily test its output.
	// We'll just call it to ensure it doesn't panic.
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("vPrintln panicked: %v", r)
		}
	}()
	vPrintln("test")
}

func TestVPrintf(t *testing.T) {
	verbose = true
	// Since this function just wraps fmt.Printf, we can't easily test its output.
	// We'll just call it to ensure it doesn't panic.
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("vPrintf panicked: %v", r)
		}
	}()
	vPrintf("test %s", "value")
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
	if err := pPng.Encode(); err == nil {
		t.Error("Portrait.Encode() for png with invalid path should have failed")
	}

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
	if err := pGif.Encode(); err == nil {
		t.Error("Portrait.Encode() for gif with invalid path should have failed")
	}

	// Test invalid style
	optsInvalidStyle := Options{
		Style: "invalid",
	}
	pInvalidStyle := NewPortrait(optsInvalidStyle)
	if err := pInvalidStyle.encodePng(); err == nil {
		t.Error("encodePng with invalid style should have failed")
	}

	// Test invalid theme
	optsInvalidTheme := Options{
		Style: "basic",
		Theme: "invalid",
	}
	pInvalidTheme := NewPortrait(optsInvalidTheme)
	if err := pInvalidTheme.encodePng(); err == nil {
		t.Error("encodePng with invalid theme should have failed")
	}
}
