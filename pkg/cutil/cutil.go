package cutil

import (
	"fmt"
	"image/color"
)

func HexToColor(s string) (*color.RGBA, error) {
	c := &color.RGBA{}

	if len(s) == 7 {
		if _, err := fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B); err != nil {
			return nil, err
		}
		c.A = 255
	}

	return c, nil
}
