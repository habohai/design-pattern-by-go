package draw

import "fmt"

type BrushPen interface {
	DrawPicture()
}

type BigBrushPen struct {
	Color
}

func (bbp BigBrushPen) DrawPicture() {
	fmt.Println("Draw picture with big brush pen")
	bbp.Use()
}

type SmallBrushPen struct {
	Color
}

func (sbp SmallBrushPen) DrawPicture() {
	fmt.Println("Draw picture with small brush pen")
	sbp.Use()
}

func NewBrushPen(t string, color Color) BrushPen {
	switch t {
	case "BIG":
		return BigBrushPen{
			Color: color,
		}
	case "SMALL":
		return SmallBrushPen{
			Color: color,
		}
	default:
		return nil
	}
}
