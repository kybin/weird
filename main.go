package main

import (
	"image"
	"image/color"

	"github.com/kybin/yeird/gui"
)

func main() {
	size := image.Pt(1280, 720)
	win := gui.NewWindow("Title", size, &gui.Area{
		Holder:  gui.Filler{},
		BgColor: &color.RGBA{0, 0, 255, 255},
		Children: []*gui.Area{
			&gui.Area{
				Name:    "header",
				Holder:  gui.TopHolder{40},
				BgColor: &color.RGBA{255, 255, 255, 255},
				Children: []*gui.Area{
					&gui.Area{
						Name:         "menus",
						Holder:       gui.Padder{5},
						BgColor:      &color.RGBA{70, 70, 90, 255},
						BorderRadius: 5,
						Children: []*gui.Area{
							&gui.Area{
								Name:   "file",
								Holder: gui.LeftHolder{100},
							},
						},
					},
				},
			},
			&gui.Area{
				Name:    "footer",
				Holder:  gui.BottomHolder{100},
				BgColor: &color.RGBA{0, 255, 0, 255},
			},
			&gui.Area{
				Name:   "body",
				Holder: gui.Filler{},
				Children: []*gui.Area{
					&gui.Area{
						Name:    "left",
						Holder:  gui.LeftHolder{300},
						BgColor: &color.RGBA{255, 0, 0, 255},
					},
					&gui.Area{
						Name:         "pad",
						Holder:       gui.Padder{20},
						BgColor:      &color.RGBA{0, 255, 255, 255},
						BorderRadius: 10,
					},
				},
			},
		},
	})
	win.Open()
}
