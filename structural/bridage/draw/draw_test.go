package draw

import (
	"fmt"
	"testing"
)

func TestDraw(t *testing.T) {
	var tColor Color
	tColor = Red{}
	tBrushPen := NewBrushPen("BIG", tColor)
	tBrushPen.DrawPicture()
	fmt.Println("=========================")
	tColor = Green{}
	tBrushPen = NewBrushPen("SMALL", tColor)
	tBrushPen.DrawPicture()
	fmt.Println("=========================")
	tColor = Yellow{}
	tBrushPen = NewBrushPen("BIG", tColor)
	tBrushPen.DrawPicture()
}
