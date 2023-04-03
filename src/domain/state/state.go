package state

import (
	"fmt"
	"io"
)

func (s State) Write(wr io.Writer, separator bool) {

	if separator {
		fmt.Fprintf(wr, "----------------------------------------\n")
	}

	fmt.Fprintf(wr, "Initial: %s\n", s.InitialNode.Name)

	fmt.Fprintf(wr, "Current: %s\n", s.Current.Name)

	fmt.Fprintf(wr, "Missing: ")

	for key := range s.MissingNodes {
		fmt.Fprintf(wr, "%s ", key)
	}

	fmt.Fprintf(wr, "\nPath: ")
	for i := 0; i < len(s.Path.Nodes)-1; i++ {
		fmt.Fprintf(wr, "%s -> ", s.Path.Nodes[i].Name)
	}

	fmt.Fprintf(wr, "%s\n", s.Path.Nodes[len(s.Path.Nodes)-1].Name)

	fmt.Fprintf(wr, "Cost: %f\n", s.Path.Cost)

	if separator {
		fmt.Fprintf(wr, "----------------------------------------")
	}

}
