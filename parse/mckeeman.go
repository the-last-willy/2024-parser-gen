package parse

func MckeemanGrammar() *Tree {
	return &Tree{Type: "grammar", Children: []*Tree{
		{Type: "rules", Children: []*Tree{
			{Type: "rule", Children: []*Tree{
				{Type: "name", Data: "grammar"},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "rules"},
					}},
				}},
			}},
			{Type: "rule", Children: []*Tree{
				{Type: "name", Data: "space"},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "codepoint", Data: "0020"},
					}},
				}},
			}},
			{Type: "rule", Children: []*Tree{
				{Type: "name", Data: "newline"},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "codepoint", Data: "000A"},
					}},
				}},
			}},
			{Type: "rule", Children: []*Tree{
				{Type: "name", Data: "name"},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "letter"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "name"},
					}},
				}},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "letter"},
					}},
				}},
			}},
			{Type: "rule", Children: []*Tree{
				{Type: "name", Data: "letter"},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "range", Children: []*Tree{
							{Type: "codepoint", Data: "a"},
							{Type: "codepoint", Data: "z"},
						}},
					}},
				}},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "range", Children: []*Tree{
							{Type: "codepoint", Data: "A"},
							{Type: "codepoint", Data: "Z"},
						}},
					}},
				}},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "codepoint", Data: "_"},
					}},
				}},
			}},
			{Type: "rule", Children: []*Tree{
				{Type: "name", Data: "indentation"},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "space"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "space"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "space"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "space"},
					}},
				}},
			}},
			{Type: "rule", Children: []*Tree{
				{Type: "name", Data: "rules"},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "rule"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "newline"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "rules"},
					}},
				}},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "rule"},
					}},
				}},
			}},
			{Type: "rule", Children: []*Tree{
				{Type: "name", Data: "rule"},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "name"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "newline"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "alternatives"},
					}},
				}},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "nothing"},
					}},
				}},
			}},
			{Type: "rule", Children: []*Tree{
				{Type: "name", Data: "nothing"},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "indentation"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "codepoint", Data: "\""},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "codepoint", Data: "\""},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "newline"},
					}},
				}},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "characters", Data: ""},
					}},
				}},
			}},
			{Type: "rule", Children: []*Tree{
				{Type: "name", Data: "alternatives"},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "alternative"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "alternatives"},
					}},
				}},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "alternative"},
					}},
				}},
			}},
			{Type: "rule", Children: []*Tree{
				{Type: "name", Data: "alternative"},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "indentation"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "items"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "newline"},
					}},
				}},
			}},
			{Type: "rule", Children: []*Tree{
				{Type: "name", Data: "items"},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "item"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "space"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "items"},
					}},
				}},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "item"},
					}},
				}},
			}},
			{Type: "rule", Children: []*Tree{
				{Type: "name", Data: "item"},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "literal"},
					}},
				}},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "name"},
					}},
				}},
			}},
			{Type: "rule", Children: []*Tree{
				{Type: "name", Data: "literal"},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "range"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "exclude"},
					}},
				}},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "singleton"},
					}},
				}},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "codepoint", Data: "\""},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "characters"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "codepoint", Data: "\""},
					}},
				}},
			}},
			{Type: "rule", Children: []*Tree{
				{Type: "name", Data: "singleton"},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "codepoint", Data: "'"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "codepoint"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "codepoint", Data: "'"},
					}},
				}},
			}},
			{Type: "rule", Children: []*Tree{
				{Type: "name", Data: "codepoint"},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "hexcode"},
					}},
				}},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "range", Children: []*Tree{
							{Type: "codepoint", Data: "0020"},
							{Type: "codepoint", Data: "10FFFF"},
						}},
					}},
				}},
			}},
			{Type: "rule", Children: []*Tree{
				{Type: "name", Data: "hexcode"},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "characters", Children: []*Tree{
							{Type: "character", Data: "1"},
							{Type: "character", Data: "0"},
						}},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "hex"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "hex"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "hex"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "hex"},
					}},
				}},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "hex"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "hex"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "hex"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "hex"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "hex"},
					}},
				}},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "hex"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "hex"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "hex"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "hex"},
					}},
				}},
			}},
			{Type: "rule", Children: []*Tree{
				{Type: "name", Data: "hex"},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "range", Children: []*Tree{
							{Type: "codepoint", Data: "0"},
							{Type: "codepoint", Data: "9"},
						}},
					}},
				}},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "range", Children: []*Tree{
							{Type: "codepoint", Data: "A"},
							{Type: "codepoint", Data: "F"},
						}},
					}},
				}},
			}},
			{Type: "rule", Children: []*Tree{
				{Type: "name", Data: "range"},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "singleton"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "space"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "codepoint", Data: "."},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "space"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "singleton"},
					}},
				}},
			}},
			{Type: "rule", Children: []*Tree{
				{Type: "name", Data: "exclude"},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "space"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "codepoint", Data: "-"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "space"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "singleton"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "exclude"},
					}},
				}},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "space"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "codepoint", Data: "-"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "space"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "range"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "exclude"},
					}},
				}},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "characters", Data: ""},
					}},
				}},
			}},
			{Type: "rule", Children: []*Tree{
				{Type: "name", Data: "characters"},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "character"},
					}},
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "characters"},
					}},
				}},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "name", Data: "character"},
					}},
				}},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "characters", Data: ""},
					}},
				}},
			}},
			{Type: "rule", Children: []*Tree{
				{Type: "name", Data: "character"},
				{Type: "alternative", Children: []*Tree{
					{Type: "item", Children: []*Tree{
						{Type: "range", Children: []*Tree{
							{Type: "codepoint", Data: "0020"},
							{Type: "codepoint", Data: "10FFFF"},
						}},
						{Type: "exclude", Children: []*Tree{
							{Type: "codepoint", Data: "\""},
						}},
					}},
				}},
			}},
		}},
	}}
}
