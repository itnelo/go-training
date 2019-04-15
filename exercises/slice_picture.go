package exercises

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	var slice [][]uint8

	for i := 0; i < dx; i++ {
		var nested []uint8

		for j := 0; j < dy; j++ {
			nested = append(nested, uint8(dx^dy))
		}

		slice = append(slice, nested)
	}

	return slice
}

func showPic() {
	pic.Show(Pic)
}
