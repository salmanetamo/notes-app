package ui

import (
	"bytes"
	"fmt"
	"strings"
)

type Cell struct {
	Value        string
	Length       int
	IsFirstInRow bool
	IsLastInRow  bool
}

type Row []Cell

func (r Row) Build() {
	for _, c := range r {
		c.build()
	}
	DrawHorizontalLine()
}

func (c Cell) build() {
	var cellContent bytes.Buffer
	if c.IsFirstInRow {
		cellContent.WriteString("\t|")
	}
	cellContent.WriteString(" ")
	lengthAvailableForValue := c.Length - 3

	if len(c.Value) <= lengthAvailableForValue {
		lengthBlankSpaceBeforeValue := (lengthAvailableForValue - len(c.Value)) / 2
		cellContent.WriteString(strings.Repeat(" ", lengthBlankSpaceBeforeValue))
		cellContent.WriteString(c.Value)
		cellContent.WriteString(strings.Repeat(" ", lengthAvailableForValue-(lengthBlankSpaceBeforeValue+len(c.Value))))
	} else {
		cellContent.WriteString(c.Value[0:lengthAvailableForValue-3] + strings.Repeat(".", 3))
	}

	cellContent.WriteString(" ")
	cellContent.WriteString("|")
	if c.IsLastInRow {
		cellContent.WriteString("\n")
	}
	fmt.Print(&cellContent)
}

func DrawHorizontalLine() {
	fmt.Println("\t" + strings.Repeat("-", TableLength))
}
