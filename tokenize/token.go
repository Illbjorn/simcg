package tokenize

/*------------------------------------------------------------------------------
 * Token
 *----------------------------------------------------------------------------*/

type Token struct {
	Kind           Kind
	Sy, Sx, Ey, Ex int
	PosStart, PosStop int
	Source            []byte
}

func (tk Token) Value() []byte {
	if len(tk.Source) == 0 {
		return nil
	}
	return tk.Source[tk.PosStart:tk.PosStop]
}

func (tk Token) ValueS() string {
	return string(tk.Value())
}
