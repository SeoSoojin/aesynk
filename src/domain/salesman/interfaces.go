package salesman

import (
	"github.com/seosoojin/aesynk/src/domain/path"
)

type Solver interface {
	Solve() (path.Path, error)
}
