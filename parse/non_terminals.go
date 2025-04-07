package parse

import (
	"github.com/illbjorn/simcg/tokenize"
)

func Seq(kind Kind, parsers ...Combinator) Combinator {
	return func(set *tokenize.Set) NodeProducer {
		return func(yield func(*Node, error) bool) {
			var (
				snap   = set.Snap()
				branch = &Node{Kind: kind}
			)

			for leaf, err := range join(set, parsers...) {
				if err != nil {
					// !
					// Not OK
					// !
					set.Revert(snap)
					yield(nil, err)
					return
				} else {
					// OK
					Push(branch, leaf)
				}
			}
			// OK
			yield(branch, nil)
		}
	}
}

func OneOf(parsers ...Combinator) Combinator {
	return func(set *tokenize.Set) NodeProducer {
		return func(yield func(*Node, error) bool) {
			var (
				snap = set.Snap()
			)

			for n, err := range join(set, parsers...) {
				if err == nil {
					// OK
					yield(n, nil)
					return
				}
				// !
				// Not OK
				// !
				set.Revert(snap)
			}
			// !
			// Not OK
			// !
			yield(nil, ErrNoMatch)
		}
	}
}

func Lazy(kind Kind, parsers ...Combinator) Combinator {
	return func(set *tokenize.Set) NodeProducer {
		return func(yield func(*Node, error) bool) {
			for {
				var (
					snap   = set.Snap()
					branch = &Node{Kind: kind}
				)

				for leaf, err := range join(set, parsers...) {
					if err != nil {
						set.Revert(snap)
						return
					}
					Push(branch, leaf)
				}

				if !yield(branch, nil) {
					return
				}
			}
		}
	}
}

func Eager(kind Kind, parsers ...Combinator) Combinator {
	return func(set *tokenize.Set) NodeProducer {
		return func(yield func(*Node, error) bool) {
			count := 0
			for {
				var (
					snap   = set.Snap()
					branch = &Node{Kind: kind}
				)

				for leaf, err := range join(set, parsers...) {
					switch {
					case err != nil && count > 0:
						set.Revert(snap)
						return

					case err != nil:
						set.Revert(snap)
						yield(nil, err)
						return

					default:
						Push(branch, leaf)
					}
				}

				if !yield(branch, nil) {
					return
				}
				count += 1
			}
		}
	}
}

func Maybe(parsers ...Combinator) Combinator {
	return func(set *tokenize.Set) NodeProducer {
		return func(yield func(*Node, error) bool) {
			var (
				snap   = set.Snap()
				branch = &Node{}
			)

			for leaf, err := range join(set, parsers...) {
				if err != nil {
					// !
					// Not OK
					// !
					set.Revert(snap)
					return
				}
				Push(branch, leaf)
			}

			// OK
			yield(branch, nil)
		}
	}
}
