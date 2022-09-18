package kata

import "fmt"

func roundRGB(v int) int {
	switch {
	case v < 0:
		return 0
	case v > 255:
		return 255
	default:
		return v
	}
}

func RGB(r, g, b int) string {
	return fmt.Sprintf("%02X%02X%02X", roundRGB(r), roundRGB(g), roundRGB(b))
}
