package parser

import "fmt"

type linkNode struct {
	Href string
	Text string
}

func (n linkNode) String() string {
	return fmt.Sprintf(
		"{\n\tHref: %q,\n\tText: %q,\n},",
		n.Href,
		n.Text,
	)
}
