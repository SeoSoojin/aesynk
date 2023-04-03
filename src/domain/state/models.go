package state

import (
	"github.com/seosoojin/aesynk/src/domain/node"
	"github.com/seosoojin/aesynk/src/domain/path"
)

type State struct {
	InitialNode  *node.Node
	Current      *node.Node
	MissingNodes map[string]struct{}
	Path         path.Path
}
