package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type ParseSource struct {
	widget.TextGrid

	highlightStyle *widget.CustomTextGridStyle
}

func NewParseSource() *ParseSource {
	ps := &ParseSource{}
	ps.ExtendBaseWidget(ps)

	ps.highlightStyle = &widget.CustomTextGridStyle{
		TextStyle: fyne.TextStyle{},
		FGColor:   nil,
		BGColor: color.RGBA{
			R: 255,
			G: 255,
			B: 0,
			A: 255,
		},
	}

	return ps
}

func (ps *ParseSource) Highlight(first, last int) {
	for _, r := range ps.Rows {
		for i, _ := range r.Cells {
			r.Cells[i].Style = &widget.CustomTextGridStyle{}
		}
	}

	idx := 0
	for i := range ps.Rows {
		for j := range ps.Rows[i].Cells {
			if idx >= first && idx < last {
				ps.Rows[i].Cells[j].Style = ps.highlightStyle
			}
			idx += 1
		}
		idx += 1
	}
}
