package generator

import (
	"errors"
	"fmt"

	"github.com/fogleman/gg"
)

var (
	ErrFileNotFound = errors.New("file not found")
	ErrRefNotFound  = errors.New("ref not found")
)

type Node struct {
	Type          string  `yaml:"type"`
	Title         string  `yaml:"title"`
	Children      []*Node `yaml:"children"`
	ScalingFactor float64 `yaml:"scale"`
	Id            string  `yaml:"id"`
	Ref           string  `yaml:"ref"`

	touched bool
	x       int
	y       int
}

const padding = 1.5

func (n *Node) Draw(ctx *gg.Context, x, y int) error {
	n.touched = true
	n.x = x
	n.y = y

	img, err := gg.LoadImage("images/" + n.Type + ".png")

	if err != nil {
		return fmt.Errorf("%s: %w", n.Type, ErrFileNotFound)
	}

	height := img.Bounds().Max.Y
	halfHeight := height / 2
	heightWithPadding := (1 + padding) * float64(height)

	ctx.DrawImageAnchored(img, x, y, 0, 0)
	ctx.DrawStringAnchored(n.Title, float64(x+halfHeight), float64(y)+float64(height)*1.1, 0.5, 0.2)

	for i, child := range n.Children {

		factor := scalingFactor(child.ScalingFactor)

		verticalOffset := float64(i*(len(child.Children)/2+1)) * heightWithPadding * factor
		horizontalOffset := int(heightWithPadding)
		arrowDestx := destCoord(child.x, x)
		arrowDesty := destCoord(child.y, y)

		drawArrow(ctx, float64(x+height), float64(y+halfHeight), float64(arrowDestx+horizontalOffset), verticalOffset+float64(arrowDesty+halfHeight))

		if !child.touched {
			err = child.Draw(ctx, x+horizontalOffset, y+int(verticalOffset))
			if err != nil {
				return ErrFileNotFound
			}
		}

	}

	return nil
}

func scalingFactor(n float64) float64 {
	if n == 0 {
		return 1
	}
	return n
}

func destCoord(n int, def int) int {
	if n == 0 {
		return def
	}
	return n
}

func (n *Node) AddId(idMap map[string]*Node) {
	if n.Id != "" {
		idMap[n.Id] = n
	}

	for _, c := range n.Children {
		c.AddId(idMap)
	}
}

func (n *Node) AddRef(idMap map[string]*Node) error {
	if n.Ref != "" {
		id, ok := idMap[n.Ref]
		if !ok {
			return ErrRefNotFound
		}

		n.Children = append(n.Children, id)
	}

	for _, c := range n.Children {
		c.AddRef(idMap)
	}

	return nil
}
