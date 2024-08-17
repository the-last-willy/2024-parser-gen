package parse

type Formatter struct {
}

func (f *Formatter) FormatGrammar(dt DataTree) string {
	if dt.Root.Type != "grammar" {
		panic("Formatter.FormatGrammar: not a grammar")
	}

	rules := dt.Root.FindAll("rule")

	str := ""

	for _, rule := range rules {
		str += f.FormatRule(dt.WithRoot(rule))
	}

	return str
}

func (f *Formatter) FormatRule(dt DataTree) string {
	if dt.Root.Type != "rule" {
		panic("Formatter.FormatRule: not a rule")
	}

	name := dt.Root.FindFirst("name")
	alternatives := dt.Root.FindAll("alternative")

	str := dt.DataOf(name) + "\n"

	for _, alternative := range alternatives {
		str += f.FormatAlternative(dt.WithRoot(alternative))
	}

	return str + "\n"
}

func (f *Formatter) FormatAlternative(dt DataTree) string {
	if dt.Root.Type != "alternative" {
		panic("Formatter.FormatAlternative: not an alternative")
	}

	items := dt.Root.FindAll("item")

	str := "    "
	for idx, item := range items {
		if idx > 0 {
			str += " "
		}
		str += f.FormatItem(dt.WithRoot(item))
	}
	str += "\n"

	return str
}

func (f *Formatter) FormatItem(dt DataTree) string {
	return dt.DataOf(dt.Root)
}
