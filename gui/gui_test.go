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
	win := NewWindow("Title", size)
	win.Area.SetBackgroundColor(color.RGBA{0, 0, 255, 255})
	header := win.Area.NewChild("header", TopHolder{100})
	header.SetBackgroundColor(color.RGBA{255, 255, 255, 255})
	footer := win.Area.NewChild("footer", BottomHolder{100})
	footer.SetBackgroundColor(color.RGBA{0, 255, 0, 255})
	body := win.Area.NewChild("body", Filler{})
	left := body.NewChild("left", LeftHolder{300})
	left.SetBackgroundColor(color.RGBA{255, 0, 0, 255})
	pad := body.NewChild("pad", Padder{20})
	pad.SetBackgroundColor(color.RGBA{0, 255, 255, 255})
	win.Area.Draw()
	f, err := os.Create("test.png")
	if err != nil {
		log.Fatal(err)
	}
	png.Encode(f, win.pixels)
}
