package gui

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"testing"
)

func TestGui(t *testing.T) {
	size := image.Pt(1280, 720)
	win := NewWindow("Title", size, &Area{
		Holder:  Filler{},
		bgColor: &color.RGBA{0, 0, 255, 255},
		Children: []*Area{
			&Area{
				Name:    "header",
				Holder:  TopHolder{100},
				bgColor: &color.RGBA{255, 255, 255, 255},
			},
			&Area{
				Name:    "footer",
				Holder:  BottomHolder{100},
				bgColor: &color.RGBA{0, 255, 0, 255},
			},
			&Area{
				Name:   "body",
				Holder: Filler{},
				Children: []*Area{
					&Area{
						Name:    "left",
						Holder:  LeftHolder{300},
						bgColor: &color.RGBA{255, 0, 0, 255},
					},
					&Area{
						Name:    "pad",
						Holder:  Padder{20},
						bgColor: &color.RGBA{0, 255, 255, 255},
					},
				},
			},
		},
	})
	win.Init()
	win.Fit()
	win.Area.Draw()
	f, err := os.Create("test.png")
	if err != nil {
		log.Fatal(err)
	}
	png.Encode(f, win.pixels)
}
