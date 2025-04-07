package parse

import (
	"iter"

	"github.com/illbjorn/simcg/tokenize"
)

type (
	NonTerminal  = func(combinators ...Combinator) Combinator
	NodeProducer = iter.Seq2[*Node, error]
	Combinator   = func(set *tokenize.Set) NodeProducer
)
