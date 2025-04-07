package parse

import (
	"slices"

	"github.com/illbjorn/simcg/tokenize"
)

func Terminal(kind Kind, kinds ...tokenize.Kind) Combinator {
	return func(set *tokenize.Set) NodeProducer {
		return func(yield func(*Node, error) bool) {
			next := set.Peek(1)

			// OK
			if slices.Contains(kinds, next.Kind) {
				set.Adv()
				yield(
					&Node{
						Kind:     kind,
						Sy:       next.Sy,
						Sx:       next.Sx,
						Ey:       next.Ey,
						Ex:       next.Ex,
						PosStart: next.PosStart,
						PosStop:  next.PosStop,
						Source:   next.Source,
					}, nil,
				)
				return
			}

			// !
			// Not OK
			// !
			yield(nil, ErrNoMatch)
		}
	}
}

func Discard(kinds ...tokenize.Kind) Combinator {
	return func(set *tokenize.Set) NodeProducer {
		return func(yield func(*Node, error) bool) {
			next := set.Peek(1)

			// OK
			if slices.Contains(kinds, next.Kind) {
				set.Adv()
				yield(&Node{Kind: 0}, nil)
				return
			}

			// !
			// Not OK
			// !
			yield(nil, ErrNoMatch)
		}
	}
}
