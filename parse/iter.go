package parse

import (
	"github.com/illbjorn/hades/simcg/tokenize"
)

func join(set *tokenize.Set, parsers ...Combinator) NodeProducer {
	return func(yield func(*Node, error) bool) {
		for _, parser := range parsers {
			for leaf, err := range parser(set) {
				if !yield(leaf, err) {
					return
				}
			}
		}
	}
}
