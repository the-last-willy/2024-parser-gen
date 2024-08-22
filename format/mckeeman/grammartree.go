package mckeeman

import (
	"parsium/parse"
	"parsium/tree"
)

func grammarTree() tree.Tree[parse.TreeData] {
	t := tree.NewSimpleTree[parse.TreeData]()

	r := t.NewNode(
		parse.TreeData{
			Type:  "grammar",
			First: 0,
			Last:  960,
		},
		[]tree.Node{
			t.NewNode(
				parse.TreeData{
					Type:  "rule",
					First: 0,
					Last:  18,
				},
				[]tree.Node{
					t.NewNode(
						parse.TreeData{
							Type:  "name",
							First: 0,
							Last:  7,
						},
						[]tree.Node{},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 8,
							Last:  18,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 12,
									Last:  17,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 12,
											Last:  17,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
				},
			),
			t.NewNode(
				parse.TreeData{
					Type:  "rule",
					First: 19,
					Last:  36,
				},
				[]tree.Node{
					t.NewNode(
						parse.TreeData{
							Type:  "name",
							First: 19,
							Last:  24,
						},
						[]tree.Node{},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 25,
							Last:  36,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 29,
									Last:  35,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "codepoint",
											First: 30,
											Last:  34,
										},
										[]tree.Node{
											t.NewNode(
												parse.TreeData{
													Type:  "hexcode",
													First: 30,
													Last:  34,
												},
												[]tree.Node{
													t.NewNode(
														parse.TreeData{
															Type:  "hex",
															First: 30,
															Last:  31,
														},
														[]tree.Node{},
													),
													t.NewNode(
														parse.TreeData{
															Type:  "hex",
															First: 31,
															Last:  32,
														},
														[]tree.Node{},
													),
													t.NewNode(
														parse.TreeData{
															Type:  "hex",
															First: 32,
															Last:  33,
														},
														[]tree.Node{},
													),
													t.NewNode(
														parse.TreeData{
															Type:  "hex",
															First: 33,
															Last:  34,
														},
														[]tree.Node{},
													),
												},
											),
										},
									),
								},
							),
						},
					),
				},
			),
			t.NewNode(
				parse.TreeData{
					Type:  "rule",
					First: 37,
					Last:  56,
				},
				[]tree.Node{
					t.NewNode(
						parse.TreeData{
							Type:  "name",
							First: 37,
							Last:  44,
						},
						[]tree.Node{},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 45,
							Last:  56,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 49,
									Last:  55,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "codepoint",
											First: 50,
											Last:  54,
										},
										[]tree.Node{
											t.NewNode(
												parse.TreeData{
													Type:  "hexcode",
													First: 50,
													Last:  54,
												},
												[]tree.Node{
													t.NewNode(
														parse.TreeData{
															Type:  "hex",
															First: 50,
															Last:  51,
														},
														[]tree.Node{},
													),
													t.NewNode(
														parse.TreeData{
															Type:  "hex",
															First: 51,
															Last:  52,
														},
														[]tree.Node{},
													),
													t.NewNode(
														parse.TreeData{
															Type:  "hex",
															First: 52,
															Last:  53,
														},
														[]tree.Node{},
													),
													t.NewNode(
														parse.TreeData{
															Type:  "hex",
															First: 53,
															Last:  54,
														},
														[]tree.Node{},
													),
												},
											),
										},
									),
								},
							),
						},
					),
				},
			),
			t.NewNode(
				parse.TreeData{
					Type:  "rule",
					First: 57,
					Last:  89,
				},
				[]tree.Node{
					t.NewNode(
						parse.TreeData{
							Type:  "name",
							First: 57,
							Last:  61,
						},
						[]tree.Node{},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 62,
							Last:  78,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 66,
									Last:  72,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 66,
											Last:  72,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 73,
									Last:  77,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 73,
											Last:  77,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 78,
							Last:  89,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 82,
									Last:  88,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 82,
											Last:  88,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
				},
			),
			t.NewNode(
				parse.TreeData{
					Type:  "rule",
					First: 90,
					Last:  133,
				},
				[]tree.Node{
					t.NewNode(
						parse.TreeData{
							Type:  "name",
							First: 90,
							Last:  96,
						},
						[]tree.Node{},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 97,
							Last:  111,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 101,
									Last:  110,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "range",
											First: 101,
											Last:  110,
										},
										[]tree.Node{
											t.NewNode(
												parse.TreeData{
													Type:  "codepoint",
													First: 102,
													Last:  103,
												},
												[]tree.Node{},
											),
											t.NewNode(
												parse.TreeData{
													Type:  "codepoint",
													First: 108,
													Last:  109,
												},
												[]tree.Node{},
											),
										},
									),
									t.NewNode(
										parse.TreeData{
											Type:  "exclude",
											First: 110,
											Last:  110,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 111,
							Last:  125,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 115,
									Last:  124,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "range",
											First: 115,
											Last:  124,
										},
										[]tree.Node{
											t.NewNode(
												parse.TreeData{
													Type:  "codepoint",
													First: 116,
													Last:  117,
												},
												[]tree.Node{},
											),
											t.NewNode(
												parse.TreeData{
													Type:  "codepoint",
													First: 122,
													Last:  123,
												},
												[]tree.Node{},
											),
										},
									),
									t.NewNode(
										parse.TreeData{
											Type:  "exclude",
											First: 124,
											Last:  124,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 125,
							Last:  133,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 129,
									Last:  132,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "codepoint",
											First: 130,
											Last:  131,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
				},
			),
			t.NewNode(
				parse.TreeData{
					Type:  "rule",
					First: 134,
					Last:  174,
				},
				[]tree.Node{
					t.NewNode(
						parse.TreeData{
							Type:  "name",
							First: 134,
							Last:  145,
						},
						[]tree.Node{},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 146,
							Last:  174,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 150,
									Last:  155,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 150,
											Last:  155,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 156,
									Last:  161,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 156,
											Last:  161,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 162,
									Last:  167,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 162,
											Last:  167,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 168,
									Last:  173,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 168,
											Last:  173,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
				},
			),
			t.NewNode(
				parse.TreeData{
					Type:  "rule",
					First: 175,
					Last:  213,
				},
				[]tree.Node{
					t.NewNode(
						parse.TreeData{
							Type:  "name",
							First: 175,
							Last:  180,
						},
						[]tree.Node{},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 181,
							Last:  204,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 185,
									Last:  189,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 185,
											Last:  189,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 190,
									Last:  197,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 190,
											Last:  197,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 198,
									Last:  203,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 198,
											Last:  203,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 204,
							Last:  213,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 208,
									Last:  212,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 208,
											Last:  212,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
				},
			),
			t.NewNode(
				parse.TreeData{
					Type:  "rule",
					First: 214,
					Last:  261,
				},
				[]tree.Node{
					t.NewNode(
						parse.TreeData{
							Type:  "name",
							First: 214,
							Last:  218,
						},
						[]tree.Node{},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 219,
							Last:  249,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 223,
									Last:  227,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 223,
											Last:  227,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 228,
									Last:  235,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 228,
											Last:  235,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 236,
									Last:  248,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 236,
											Last:  248,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 249,
							Last:  261,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 253,
									Last:  260,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 253,
											Last:  260,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
				},
			),
			t.NewNode(
				parse.TreeData{
					Type:  "rule",
					First: 262,
					Last:  309,
				},
				[]tree.Node{
					t.NewNode(
						parse.TreeData{
							Type:  "name",
							First: 262,
							Last:  269,
						},
						[]tree.Node{},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 270,
							Last:  302,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 274,
									Last:  285,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 274,
											Last:  285,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 286,
									Last:  289,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "codepoint",
											First: 287,
											Last:  288,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 290,
									Last:  293,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "codepoint",
											First: 291,
											Last:  292,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 294,
									Last:  301,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 294,
											Last:  301,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 302,
							Last:  309,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 306,
									Last:  308,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "characters",
											First: 307,
											Last:  307,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
				},
			),
			t.NewNode(
				parse.TreeData{
					Type:  "rule",
					First: 310,
					Last:  368,
				},
				[]tree.Node{
					t.NewNode(
						parse.TreeData{
							Type:  "name",
							First: 310,
							Last:  322,
						},
						[]tree.Node{},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 323,
							Last:  352,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 327,
									Last:  338,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 327,
											Last:  338,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 339,
									Last:  351,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 339,
											Last:  351,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 352,
							Last:  368,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 356,
									Last:  367,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 356,
											Last:  367,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
				},
			),
			t.NewNode(
				parse.TreeData{
					Type:  "rule",
					First: 369,
					Last:  411,
				},
				[]tree.Node{
					t.NewNode(
						parse.TreeData{
							Type:  "name",
							First: 369,
							Last:  380,
						},
						[]tree.Node{},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 381,
							Last:  411,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 385,
									Last:  396,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 385,
											Last:  396,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 397,
									Last:  402,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 397,
											Last:  402,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 403,
									Last:  410,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 403,
											Last:  410,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
				},
			),
			t.NewNode(
				parse.TreeData{
					Type:  "rule",
					First: 412,
					Last:  448,
				},
				[]tree.Node{
					t.NewNode(
						parse.TreeData{
							Type:  "name",
							First: 412,
							Last:  417,
						},
						[]tree.Node{},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 418,
							Last:  439,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 422,
									Last:  426,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 422,
											Last:  426,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 427,
									Last:  432,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 427,
											Last:  432,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 433,
									Last:  438,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 433,
											Last:  438,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 439,
							Last:  448,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 443,
									Last:  447,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 443,
											Last:  447,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
				},
			),
			t.NewNode(
				parse.TreeData{
					Type:  "rule",
					First: 449,
					Last:  475,
				},
				[]tree.Node{
					t.NewNode(
						parse.TreeData{
							Type:  "name",
							First: 449,
							Last:  453,
						},
						[]tree.Node{},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 454,
							Last:  466,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 458,
									Last:  465,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 458,
											Last:  465,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 466,
							Last:  475,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 470,
									Last:  474,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 470,
											Last:  474,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
				},
			),
			t.NewNode(
				parse.TreeData{
					Type:  "rule",
					First: 476,
					Last:  539,
				},
				[]tree.Node{
					t.NewNode(
						parse.TreeData{
							Type:  "name",
							First: 476,
							Last:  483,
						},
						[]tree.Node{},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 484,
							Last:  502,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 488,
									Last:  493,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 488,
											Last:  493,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 494,
									Last:  501,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 494,
											Last:  501,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 502,
							Last:  516,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 506,
									Last:  515,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 506,
											Last:  515,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 516,
							Last:  539,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 520,
									Last:  523,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "codepoint",
											First: 521,
											Last:  522,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 524,
									Last:  534,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 524,
											Last:  534,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 535,
									Last:  538,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "codepoint",
											First: 536,
											Last:  537,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
				},
			),
			t.NewNode(
				parse.TreeData{
					Type:  "rule",
					First: 540,
					Last:  572,
				},
				[]tree.Node{
					t.NewNode(
						parse.TreeData{
							Type:  "name",
							First: 540,
							Last:  549,
						},
						[]tree.Node{},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 550,
							Last:  572,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 554,
									Last:  557,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "codepoint",
											First: 555,
											Last:  556,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 558,
									Last:  567,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 558,
											Last:  567,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 568,
									Last:  571,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "codepoint",
											First: 569,
											Last:  570,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
				},
			),
			t.NewNode(
				parse.TreeData{
					Type:  "rule",
					First: 573,
					Last:  617,
				},
				[]tree.Node{
					t.NewNode(
						parse.TreeData{
							Type:  "name",
							First: 573,
							Last:  582,
						},
						[]tree.Node{},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 583,
							Last:  595,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 587,
									Last:  594,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 587,
											Last:  594,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 595,
							Last:  617,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 599,
									Last:  616,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "range",
											First: 599,
											Last:  616,
										},
										[]tree.Node{
											t.NewNode(
												parse.TreeData{
													Type:  "codepoint",
													First: 600,
													Last:  604,
												},
												[]tree.Node{
													t.NewNode(
														parse.TreeData{
															Type:  "hexcode",
															First: 600,
															Last:  604,
														},
														[]tree.Node{
															t.NewNode(
																parse.TreeData{
																	Type:  "hex",
																	First: 600,
																	Last:  601,
																},
																[]tree.Node{},
															),
															t.NewNode(
																parse.TreeData{
																	Type:  "hex",
																	First: 601,
																	Last:  602,
																},
																[]tree.Node{},
															),
															t.NewNode(
																parse.TreeData{
																	Type:  "hex",
																	First: 602,
																	Last:  603,
																},
																[]tree.Node{},
															),
															t.NewNode(
																parse.TreeData{
																	Type:  "hex",
																	First: 603,
																	Last:  604,
																},
																[]tree.Node{},
															),
														},
													),
												},
											),
											t.NewNode(
												parse.TreeData{
													Type:  "codepoint",
													First: 609,
													Last:  615,
												},
												[]tree.Node{
													t.NewNode(
														parse.TreeData{
															Type:  "hexcode",
															First: 609,
															Last:  615,
														},
														[]tree.Node{
															t.NewNode(
																parse.TreeData{
																	Type:  "hex",
																	First: 611,
																	Last:  612,
																},
																[]tree.Node{},
															),
															t.NewNode(
																parse.TreeData{
																	Type:  "hex",
																	First: 612,
																	Last:  613,
																},
																[]tree.Node{},
															),
															t.NewNode(
																parse.TreeData{
																	Type:  "hex",
																	First: 613,
																	Last:  614,
																},
																[]tree.Node{},
															),
															t.NewNode(
																parse.TreeData{
																	Type:  "hex",
																	First: 614,
																	Last:  615,
																},
																[]tree.Node{},
															),
														},
													),
												},
											),
										},
									),
									t.NewNode(
										parse.TreeData{
											Type:  "exclude",
											First: 616,
											Last:  616,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
				},
			),
			t.NewNode(
				parse.TreeData{
					Type:  "rule",
					First: 618,
					Last:  695,
				},
				[]tree.Node{
					t.NewNode(
						parse.TreeData{
							Type:  "name",
							First: 618,
							Last:  625,
						},
						[]tree.Node{},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 626,
							Last:  651,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 630,
									Last:  634,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "characters",
											First: 631,
											Last:  633,
										},
										[]tree.Node{
											t.NewNode(
												parse.TreeData{
													Type:  "character",
													First: 631,
													Last:  632,
												},
												[]tree.Node{},
											),
											t.NewNode(
												parse.TreeData{
													Type:  "character",
													First: 632,
													Last:  633,
												},
												[]tree.Node{},
											),
										},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 635,
									Last:  638,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 635,
											Last:  638,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 639,
									Last:  642,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 639,
											Last:  642,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 643,
									Last:  646,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 643,
											Last:  646,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 647,
									Last:  650,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 647,
											Last:  650,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 651,
							Last:  675,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 655,
									Last:  658,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 655,
											Last:  658,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 659,
									Last:  662,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 659,
											Last:  662,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 663,
									Last:  666,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 663,
											Last:  666,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 667,
									Last:  670,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 667,
											Last:  670,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 671,
									Last:  674,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 671,
											Last:  674,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 675,
							Last:  695,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 679,
									Last:  682,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 679,
											Last:  682,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 683,
									Last:  686,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 683,
											Last:  686,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 687,
									Last:  690,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 687,
											Last:  690,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 691,
									Last:  694,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 691,
											Last:  694,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
				},
			),
			t.NewNode(
				parse.TreeData{
					Type:  "rule",
					First: 696,
					Last:  728,
				},
				[]tree.Node{
					t.NewNode(
						parse.TreeData{
							Type:  "name",
							First: 696,
							Last:  699,
						},
						[]tree.Node{},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 700,
							Last:  714,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 704,
									Last:  713,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "range",
											First: 704,
											Last:  713,
										},
										[]tree.Node{
											t.NewNode(
												parse.TreeData{
													Type:  "codepoint",
													First: 705,
													Last:  706,
												},
												[]tree.Node{},
											),
											t.NewNode(
												parse.TreeData{
													Type:  "codepoint",
													First: 711,
													Last:  712,
												},
												[]tree.Node{},
											),
										},
									),
									t.NewNode(
										parse.TreeData{
											Type:  "exclude",
											First: 713,
											Last:  713,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 714,
							Last:  728,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 718,
									Last:  727,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "range",
											First: 718,
											Last:  727,
										},
										[]tree.Node{
											t.NewNode(
												parse.TreeData{
													Type:  "codepoint",
													First: 719,
													Last:  720,
												},
												[]tree.Node{},
											),
											t.NewNode(
												parse.TreeData{
													Type:  "codepoint",
													First: 725,
													Last:  726,
												},
												[]tree.Node{},
											),
										},
									),
									t.NewNode(
										parse.TreeData{
											Type:  "exclude",
											First: 727,
											Last:  727,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
				},
			),
			t.NewNode(
				parse.TreeData{
					Type:  "rule",
					First: 729,
					Last:  775,
				},
				[]tree.Node{
					t.NewNode(
						parse.TreeData{
							Type:  "name",
							First: 729,
							Last:  734,
						},
						[]tree.Node{},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 735,
							Last:  775,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 739,
									Last:  748,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 739,
											Last:  748,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 749,
									Last:  754,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 749,
											Last:  754,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 755,
									Last:  758,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "codepoint",
											First: 756,
											Last:  757,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 759,
									Last:  764,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 759,
											Last:  764,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 765,
									Last:  774,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 765,
											Last:  774,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
				},
			),
			t.NewNode(
				parse.TreeData{
					Type:  "rule",
					First: 776,
					Last:  863,
				},
				[]tree.Node{
					t.NewNode(
						parse.TreeData{
							Type:  "name",
							First: 776,
							Last:  783,
						},
						[]tree.Node{},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 784,
							Last:  822,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 788,
									Last:  793,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 788,
											Last:  793,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 794,
									Last:  797,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "codepoint",
											First: 795,
											Last:  796,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 798,
									Last:  803,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 798,
											Last:  803,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 804,
									Last:  813,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 804,
											Last:  813,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 814,
									Last:  821,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 814,
											Last:  821,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 822,
							Last:  856,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 826,
									Last:  831,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 826,
											Last:  831,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 832,
									Last:  835,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "codepoint",
											First: 833,
											Last:  834,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 836,
									Last:  841,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 836,
											Last:  841,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 842,
									Last:  847,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 842,
											Last:  847,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 848,
									Last:  855,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 848,
											Last:  855,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 856,
							Last:  863,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 860,
									Last:  862,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "characters",
											First: 861,
											Last:  861,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
				},
			),
			t.NewNode(
				parse.TreeData{
					Type:  "rule",
					First: 864,
					Last:  921,
				},
				[]tree.Node{
					t.NewNode(
						parse.TreeData{
							Type:  "name",
							First: 864,
							Last:  874,
						},
						[]tree.Node{},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 875,
							Last:  900,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 879,
									Last:  888,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 879,
											Last:  888,
										},
										[]tree.Node{},
									),
								},
							),
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 889,
									Last:  899,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 889,
											Last:  899,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 900,
							Last:  914,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 904,
									Last:  913,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "name",
											First: 904,
											Last:  913,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 914,
							Last:  921,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 918,
									Last:  920,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "characters",
											First: 919,
											Last:  919,
										},
										[]tree.Node{},
									),
								},
							),
						},
					),
				},
			),
			t.NewNode(
				parse.TreeData{
					Type:  "rule",
					First: 922,
					Last:  960,
				},
				[]tree.Node{
					t.NewNode(
						parse.TreeData{
							Type:  "name",
							First: 922,
							Last:  931,
						},
						[]tree.Node{},
					),
					t.NewNode(
						parse.TreeData{
							Type:  "alternative",
							First: 932,
							Last:  960,
						},
						[]tree.Node{
							t.NewNode(
								parse.TreeData{
									Type:  "item",
									First: 936,
									Last:  959,
								},
								[]tree.Node{
									t.NewNode(
										parse.TreeData{
											Type:  "range",
											First: 936,
											Last:  953,
										},
										[]tree.Node{
											t.NewNode(
												parse.TreeData{
													Type:  "codepoint",
													First: 937,
													Last:  941,
												},
												[]tree.Node{
													t.NewNode(
														parse.TreeData{
															Type:  "hexcode",
															First: 937,
															Last:  941,
														},
														[]tree.Node{
															t.NewNode(
																parse.TreeData{
																	Type:  "hex",
																	First: 937,
																	Last:  938,
																},
																[]tree.Node{},
															),
															t.NewNode(
																parse.TreeData{
																	Type:  "hex",
																	First: 938,
																	Last:  939,
																},
																[]tree.Node{},
															),
															t.NewNode(
																parse.TreeData{
																	Type:  "hex",
																	First: 939,
																	Last:  940,
																},
																[]tree.Node{},
															),
															t.NewNode(
																parse.TreeData{
																	Type:  "hex",
																	First: 940,
																	Last:  941,
																},
																[]tree.Node{},
															),
														},
													),
												},
											),
											t.NewNode(
												parse.TreeData{
													Type:  "codepoint",
													First: 946,
													Last:  952,
												},
												[]tree.Node{
													t.NewNode(
														parse.TreeData{
															Type:  "hexcode",
															First: 946,
															Last:  952,
														},
														[]tree.Node{
															t.NewNode(
																parse.TreeData{
																	Type:  "hex",
																	First: 948,
																	Last:  949,
																},
																[]tree.Node{},
															),
															t.NewNode(
																parse.TreeData{
																	Type:  "hex",
																	First: 949,
																	Last:  950,
																},
																[]tree.Node{},
															),
															t.NewNode(
																parse.TreeData{
																	Type:  "hex",
																	First: 950,
																	Last:  951,
																},
																[]tree.Node{},
															),
															t.NewNode(
																parse.TreeData{
																	Type:  "hex",
																	First: 951,
																	Last:  952,
																},
																[]tree.Node{},
															),
														},
													),
												},
											),
										},
									),
									t.NewNode(
										parse.TreeData{
											Type:  "exclude",
											First: 953,
											Last:  959,
										},
										[]tree.Node{
											t.NewNode(
												parse.TreeData{
													Type:  "codepoint",
													First: 957,
													Last:  958,
												},
												[]tree.Node{},
											),
										},
									),
								},
							),
						},
					),
				},
			),
		},
	)

	return t.WithRoot(r)
}
