package gui

import (
	"fmt"
	"image"
	"image/color"
	"log"

	"golang.org/x/exp/shiny/driver"
	"golang.org/x/exp/shiny/screen"
	"golang.org/x/mobile/event/key"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/mouse"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
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
	Name  string
	Full  image.Rectangle
	Avail image.Rectangle

	Window *Window
	Parent *Area
	Holder PlaceHolder

	bgColor *color.RGBA

	ChildMap map[string]*Area
	Children []*Area
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

func (a *Area) DoRecursive(f func(*Area)) {
	f(a)
	for _, ch := range a.Children {
		ch.DoRecursive(f)
	}
}

type Window struct {
	Size   image.Point
	Area   *Area
	pixels *image.RGBA
}

func NewWindow(title string, size image.Point, area *Area) *Window {
	win := &Window{
		Size: size,
		Area: area,
	}
	return win
}

func (win *Window) Open() {
	win.Init()
	win.Fit()
	driver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(&screen.NewWindowOptions{Width: win.Size.X, Height: win.Size.Y})
		if err != nil {
			log.Fatal(err)
		}
		defer w.Release()

		buf, err := s.NewBuffer(win.Size)
		if err != nil {
			log.Fatal(err)
		}
		win.pixels = buf.RGBA()
		for {
			switch e := w.NextEvent().(type) {
			case lifecycle.Event:
				if e.To == lifecycle.StageDead {
					return
				}
			case key.Event:
			case mouse.Event:
			case size.Event:
				win.Size = image.Point{e.WidthPx, e.HeightPx}
				buf, err = s.NewBuffer(win.Size)
				if err != nil {
					log.Fatal(err)
				}
				win.pixels = buf.RGBA()
				win.Fit()
				win.Draw()
			case paint.Event:
				w.Upload(image.Point{}, buf, buf.Bounds())
				w.Publish()
			}
		}
	})
}

func (w *Window) Init() {
	initAreaRecursive(w.Area, w, nil)
}

func initAreaRecursive(a *Area, w *Window, p *Area) {
	a.Window = w
	a.Parent = p
	a.ChildMap = make(map[string]*Area)
	for _, ch := range a.Children {
		_, ok := a.ChildMap[ch.Name]
		if ok {
			panic(fmt.Sprintf("already have child with name: %v", ch.Name))
		}
		a.ChildMap[ch.Name] = ch
		initAreaRecursive(ch, w, a)
	}
}

func (w *Window) Fit() {
	w.Area.Full = image.Rectangle{image.Point{}, w.Size}
	w.Area.Avail = w.Area.Full

	w.Area.DoRecursive(func(a *Area) {
		if a.Parent == nil { // w.Area
			return
		}
		hold, remain := a.Holder.Hold(a.Parent.Avail)
		a.Parent.Avail = remain
		a.Full = hold
		a.Avail = hold
	})
}

func (w *Window) Draw() {
	w.Area.DoRecursive(func(a *Area) {
		r := a.Full
		for y := r.Min.Y; y < r.Max.Y; y++ {
			for x := r.Min.X; x < r.Max.X; x++ {
				w.pixels.Set(x, y, a.BackgroundColor())
			}
		}
	})
}
