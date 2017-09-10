package gui

import "image/color"

type Point struct {
	X int
	Y int
}

func Pt(X, Y int) Point {
	return Point{X, Y}
}

type Rect struct {
	Min Point
	Max Point
}

func RectFromTwoPoints(min, max Point) Rect {
	if max.X < min.X {
		min.X, max.X = max.X, min.X
	}
	if max.Y < min.Y {
		min.Y, max.Y = max.Y, min.Y
	}
	return Rect{
		Min: min,
		Max: max,
	}
}

func (a Rect) Width() uint {
	return uint(a.Max.X - a.Min.X)
}

func (a Rect) Height() uint {
	return uint(a.Max.Y - a.Min.Y)
}

type PlaceHolder interface {
	Hold(Rect) (hold, remain Rect)
}

type TopHolder struct {
	Height uint
}

func (h TopHolder) Hold(r Rect) (Rect, Rect) {
	hold := r
	remain := r
	y := r.Min.Y
	y += int(h.Height)
	if y > r.Max.Y {
		y = r.Max.Y
	}
	hold.Max.Y = y
	remain.Min.Y = y
	return hold, remain
}

type BottomHolder struct {
	Height uint
}

func (h BottomHolder) Hold(r Rect) (Rect, Rect) {
	hold := r
	remain := r
	y := r.Max.Y
	y -= int(h.Height)
	if y < r.Min.Y {
		y = r.Min.Y
	}
	hold.Min.Y = y
	remain.Max.Y = y
	return hold, remain
}

type LeftHolder struct {
	Width uint
}

func (h LeftHolder) Hold(r Rect) (Rect, Rect) {
	hold := r
	remain := r
	x := r.Min.X
	x += int(h.Width)
	if x > r.Max.X {
		x = r.Max.X
	}
	hold.Max.X = x
	remain.Min.X = x
	return hold, remain
}

type RightHolder struct {
	Width uint
}

func (h RightHolder) Hold(r Rect) (Rect, Rect) {
	hold := r
	remain := r
	x := r.Max.X
	x -= int(h.Width)
	if x < r.Min.X {
		x = r.Min.X
	}
	hold.Min.X = x
	remain.Max.X = x
	return hold, remain
}

type Filler struct{}

func (f Filler) Hold(r Rect) (Rect, Rect) {
	return r, RectFromTwoPoints(r.Max, r.Max)
}

type Area struct {
	Full  Rect
	Avail Rect

	Window   *Window
	Parent   *Area
	Children map[string]*Area

	BackgroundColor color.RGBA
}

func NewArea(rect Rect, window *Window, parent *Area) *Area {
	return &Area{
		Full:     rect,
		Avail:    rect,
		Window:   window,
		Parent:   parent,
		Children: make(map[string]*Area),
	}
}

func (a *Area) Draw() {
	pixels := a.Window.pixels
	width := int(a.Window.Area.Full.Width())
	r := a.Full
	for y := r.Min.Y; y < r.Max.Y; y++ {
		for x := r.Min.X; x < r.Max.X; x++ {
			pixels[width*y+x] = a.BackgroundColor
		}
	}
	for _, child := range a.Children {
		child.Draw()
	}
}

func (a *Area) NewChild(name string, h PlaceHolder) *Area {
	hold, remain := h.Hold(a.Avail)
	a.Avail = remain
	child := NewArea(hold, a.Window, a)
	a.Children[name] = child
	return child
}

type Window struct {
	Area   *Area
	pixels []color.Color
}

func NewWindow(title string, size Point) *Window {
	rect := RectFromTwoPoints(Pt(0, 0), size)
	win := &Window{}
	win.Area = NewArea(rect, win, nil)
	win.pixels = make([]color.Color, size.X*size.Y)
	return win
}
