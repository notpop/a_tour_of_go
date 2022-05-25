package main

// import "tour/pic"

// import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	image := make([][]uint8, dy)
	for index:= 0; index < dy; index ++ {
		image[index] = make([]uint8, dx)
		for inner_index:= 0; inner_index < dx; inner_index ++ {
			image[index][inner_index] = uint8(inner_index * index)
		}
	}

	return image
}

func main() {
	// pic.Show(Pic)
}

