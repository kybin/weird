package gui

import "testing"

func TestGui(t *testing.T) {
	size := Pt(1280, 720)
	win := NewWindow("Title", size)
	win.Area.NewChild("header", TopHolder{100})
	win.Area.NewChild("footer", BottomHolder{100})
	body := win.Area.NewChild("body", Filler{})
	body.NewChild("left", LeftHolder{300})
	// right := body.NewChild("left", LeftHolder{300})
	// right.Set("backgroundColor", "#112233")
}
