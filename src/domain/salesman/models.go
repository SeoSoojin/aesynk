package salesman

import (
	"github.com/seosoojin/aesynk/src/domain/node"
	"github.com/seosoojin/aesynk/src/domain/path"
)

type State struct {
	Current *node.Node
	Visited map[string]struct{}
	Path    path.Path
}
