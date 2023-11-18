package generator

import (
	"errors"
	"fmt"

	"github.com/fogleman/gg"
)

var (
	ErrFileNotFound = errors.New("file not found")
)

type Node struct {
	Type          string  `yaml:"type"`
	Title         string  `yaml:"title"`
	Children      []Node  `yaml:"children"`
	ScalingFactor float64 `yaml:"scale"`
}

const padding = 1.5

func (n *Node) Draw(ctx *gg.Context, x, y int) error {

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

		factor := child.ScalingFactor
		if factor == 0 {
			factor = 1
		}

		verticalOffset := float64(i*(len(child.Children)/2+1)) * heightWithPadding * factor
		horizontalOffset := int(heightWithPadding)

		drawArrow(ctx, float64(x+height), float64(y+halfHeight), float64(x+horizontalOffset), verticalOffset+float64(y+halfHeight))
		err := child.Draw(ctx, x+horizontalOffset, y+int(verticalOffset))

		if err != nil {
			return ErrFileNotFound
		}

	}

	return nil
}
