package scratch

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/madvikinggod/AdventOfCode/2019/location"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

var images = []*image.RGBA{}
var MIN = location.New(0, 0)
var MAX = location.New(0, 0)

func Output(hull map[location.Location]int, loc location.Location, step int) {
	min := location.New(0, 0)
	max := location.New(0, 0)

	for loc := range hull {
		min = min.Min(loc)
		max = max.Max(loc)
	}
	img := image.NewRGBA(image.Rectangle{
		Min: min.Point().Add(image.Pt(-1, -1)),
		Max: max.Point().Add(image.Pt(2, 2)),
	})

	for loc, clr := range hull {
		pt := loc.Point()
		if clr == 1 {
			img.Set(pt.X, pt.Y, color.White)
		} else {
			img.Set(pt.X, pt.Y, color.Black)
		}
	}

	if hull[location.New(0, 0)] == 0 {
		img.Set(0, 0, color.RGBA{50, 50, 50, 255})
	} else {
		img.Set(0, 0, color.RGBA{200, 200, 200, 255})
	}
	img.Set(loc.Point().X, loc.Point().Y, color.RGBA{0, 0, 255, 255})

	images = append(images, img)
	MAX = MAX.Max(max)
	MIN = MIN.Min(min)
}

func WriteGif() {

	for i, img := range images {
		if i%100 != 0 && i != len(images)-1 {
			continue
		}
		f, _ := os.Create(fmt.Sprintf("scratch/output-%d.png", i))
		defer f.Close()
		tmpimg := image.NewRGBA(image.Rectangle{
			Min: MIN.Point(),
			Max: MAX.Point(),
		})
		draw.Draw(tmpimg, tmpimg.Rect, image.NewUniform(color.RGBA{10, 10, 10, 255}), tmpimg.Rect.Min, draw.Src)
		draw.Draw(tmpimg, img.Rect, img, img.Rect.Min, draw.Src)
		t := imaging.FlipV(tmpimg)
		png.Encode(f, t)

	}

}
