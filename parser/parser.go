package parser

import (
	"os"

	"github.com/gllmjln/arch-gen/generator"
	"gopkg.in/yaml.v3"
)

func ParseInput(inputLoc string) (*generator.Node, error) {
	nodes := generator.Node{}

	dat, err := os.ReadFile(inputLoc)

	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(dat, &nodes)

	if err != nil {
		return nil, err
	}

	return &nodes, nil
}
