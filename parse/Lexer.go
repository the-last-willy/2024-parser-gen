package parse

type Lexer struct {
	Grammar func(stack string) *TreeData

	Output func(token TreeData)

	Offset int
	Buffer string
}

func NewLexer() *Lexer {
	return &Lexer{}
}

func (p *Lexer) Process(symbol uint8) {
	p.shift(symbol)
	for p.reduce() {
	}
}

func (p *Lexer) Flush() {
	p.Process(0x00)
}

func (p *Lexer) shift(symbol uint8) {
	p.Buffer = p.Buffer + string(symbol)
}

func (p *Lexer) reduce() bool {
	if len(p.Buffer) == 0 {
		return false
	}

	d := p.Grammar(p.Buffer)
	if d == nil {
		return false
	}

	p.Buffer = p.Buffer[d.Last:]

	shifted := *d
	shifted.First += p.Offset
	shifted.Last += p.Offset

	p.Offset += d.Last

	p.Output(shifted)

	return true
}
