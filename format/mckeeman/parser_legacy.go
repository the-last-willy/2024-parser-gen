package mckeeman

//
//import "parsium/parse"
//
//type Parser struct {
//	internal *parse.Parser
//}
//
//func NewParser() *Parser {
//	return &Parser{
//		internal: parse.NewParser(parse.MckeemanGrammar(), ""),
//	}
//}
//
//func NewParserForGrammar(dt parse.DataTree) *Parser {
//	return &Parser{
//		internal: parse.NewParser(dt.Root, dt.Store),
//	}
//}
//
//func (p *Parser) Parse(s string) parse.DataTree {
//	return parse.DataTree{
//		Root:  p.internal.Parse(s),
//		Store: s,
//	}
//}
//
//func (p *Parser) ParseRule(name string, s string) parse.DataTree {
//	return parse.DataTree{
//		Root:  p.internal.ParseRule(p.internal.Rules[name], s, 0),
//		Store: s,
//	}
//}
