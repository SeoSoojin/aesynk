package search

import (
	"github.com/seosoojin/aesynk/src/domain/node"
	"github.com/seosoojin/aesynk/src/domain/path"
)

type Searcher interface {
	FindShortestWay(map[string]*node.Node, string, string) (path.Path, bool)
}
