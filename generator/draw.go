package generator

import "github.com/fogleman/gg"

func drawArrow(ctx *gg.Context, x1, y1, x2, y2 float64) {

	if y1 == y2 {
		ctx.DrawLine(x1, y1, x2, y2)
	} else {
		ctx.NewSubPath()
		mid := x1 + (x2-x1)/2
		ctx.DrawLine(x1, y1, mid, y1)
		ctx.DrawLine(mid, y1, mid, y2)
		ctx.DrawLine(mid, y2, x2, y2)
	}
	ctx.Stroke()

	//TODO: Add some arrow

}
