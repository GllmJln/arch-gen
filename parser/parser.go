package parser

import (
	"os"

	"github.com/gllmjln/arch-gen/generator"
	"gopkg.in/yaml.v3"
)

type Parser struct {
	nodes map[string]*generator.Node
	Root  *generator.Node
}

func (p *Parser) ParseInput(inputLoc string) error {
	nodes := generator.Node{}

	dat, err := os.ReadFile(inputLoc)

	if err != nil {
		return err
	}

	err = yaml.Unmarshal(dat, &nodes)

	if err != nil {
		return err
	}

	p.Root = &nodes

	p.createCycles()
	return nil
}

func (p *Parser) createCycles() {
	p.nodes = make(map[string]*generator.Node)

	p.Root.AddId(p.nodes)
	p.Root.AddRef(p.nodes)
}
