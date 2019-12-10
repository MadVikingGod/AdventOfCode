package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

var Width = 25
var Height = 6

func main() {
	fmt.Println(len(input) / Width / Height)

	min := 99999
	score := 0
	i := 0
	for ; i < len(input)/Width/Height; i++ {
		zero := 0
		one := 0
		two := 0
		layer := i * Width * Height
		for _, char := range input[layer : layer+Width*Height] {
			switch char {
			case '0':
				zero++
			case '1':
				one++
			case '2':
				two++
			}
		}
		if zero < min {
			min = zero
			score = one * two
			fmt.Println(min, score)
		}

	}
	fmt.Println(i)

	img := image.NewRGBA(image.Rect(0, 0, Width, Height))
	img2 := make([]int, Height*Width)
	for i := 0; i < Height*Width; i++ {
		img2[i] = 2
	}

	draw.Draw(img, image.Rect(0, 0, Width, Height), image.Transparent, image.ZP, draw.Src)

	for i := 0; i < len(input)/Width/Height; i++ {
		layer := i * Width * Height

		for j, char := range input[layer : layer+Width*Height] {
			y := j / Width
			x := j % Width

			if img2[j] != 2 {
				continue
			}

			switch char {
			case '0':
				img.Set(x, y, color.Black)
				img2[j] = 0

			case '1':
				img.Set(x, y, color.White)
				img2[j] = 1
			}

		}

	}

	f, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}

	png.Encode(f, img)
	for i := 0; i < Height; i++ {
		fmt.Println(img2[i*Width : i*Width+Width])
	}

}
