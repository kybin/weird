package gui

import (
	"image"
	"image/color"
)

type PlaceHolder interface {
	Hold(image.Rectangle) (hold, remain image.Rectangle)
}

type TopHolder struct {
	Height uint
}

func (h TopHolder) Hold(r image.Rectangle) (image.Rectangle, image.Rectangle) {
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

func (h BottomHolder) Hold(r image.Rectangle) (image.Rectangle, image.Rectangle) {
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

func (h LeftHolder) Hold(r image.Rectangle) (image.Rectangle, image.Rectangle) {
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

func (h RightHolder) Hold(r image.Rectangle) (image.Rectangle, image.Rectangle) {
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

func (f Filler) Hold(r image.Rectangle) (image.Rectangle, image.Rectangle) {
	return r, image.Rectangle{r.Max, r.Max}
}

type Padder struct {
	Pad uint
}

func (p Padder) Hold(r image.Rectangle) (image.Rectangle, image.Rectangle) {
	sx := r.Max.X - r.Min.X
	sy := r.Max.Y - r.Min.Y
	var xMin, xMax int
	if sx < 2*int(p.Pad) {
		xMin = r.Min.X + sx/2
		xMax = xMin
	} else {
		xMin = r.Min.X + int(p.Pad)
		xMax = r.Max.X - int(p.Pad)
	}
	var yMin, yMax int
	if sy < 2*int(p.Pad) {
		yMin = r.Min.Y + sy/2
		yMax = yMin
	} else {
		yMin = r.Min.Y + int(p.Pad)
		yMax = r.Max.Y - int(p.Pad)
	}
	return image.Rect(xMin, yMin, xMax, yMax), image.Rectangle{r.Max, r.Max}
}

type Area struct {
	Full  image.Rectangle
	Avail image.Rectangle

	Window   *Window
	Parent   *Area
	Children map[string]*Area

	bgColor *color.RGBA
}

func NewArea(rect image.Rectangle, window *Window, parent *Area) *Area {
	return &Area{
		Full:     rect,
		Avail:    rect,
		Window:   window,
		Parent:   parent,
		Children: make(map[string]*Area),
	}
}

func (a *Area) SetBackgroundColor(c color.RGBA) {
	a.bgColor = &c
}

func (a *Area) BackgroundColor() color.RGBA {
	p := a
	for p != nil {
		if p.bgColor != nil {
			// TODO: composite with it's parent color.
			return *p.bgColor
		}
		p = p.Parent
	}
	return color.RGBA{}
}

func (a *Area) Draw() {
	pixels := a.Window.pixels
	r := a.Full
	for y := r.Min.Y; y < r.Max.Y; y++ {
		for x := r.Min.X; x < r.Max.X; x++ {
			pixels.Set(x, y, a.BackgroundColor())
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
	pixels *image.RGBA
}

func NewWindow(title string, size image.Point) *Window {
	rect := image.Rectangle{image.Pt(0, 0), size}
	win := &Window{}
	win.Area = NewArea(rect, win, nil)
	win.pixels = image.NewRGBA(rect)
	return win
}
