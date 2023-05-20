package notes

import (
	"fmt"
	"notes/ui"
	"strconv"
	"time"
)

type Note struct {
	id          int
	title       string
	description string
	createdAt   time.Time
	updatedAt   time.Time
}

func (n *Note) SetTitle(title string) {
	n.title = title
}

func (n *Note) SetDescription(description string) {
	n.description = description
}

func (n Note) GetRow() ui.Row {
	return ui.Row{
		ui.Cell{Value: strconv.Itoa(n.id), Length: ui.IdColumnLength, IsFirstInRow: true, IsLastInRow: false},
		ui.Cell{Value: n.title, Length: ui.TitleColumnLength, IsFirstInRow: false, IsLastInRow: false},
		ui.Cell{Value: n.description, Length: ui.DescriptionColumnLength, IsFirstInRow: false, IsLastInRow: false},
		ui.Cell{Value: n.createdAt.String(), Length: ui.DateColumnsLength, IsFirstInRow: false, IsLastInRow: false},
		ui.Cell{Value: n.updatedAt.String(), Length: ui.DateColumnsLength, IsFirstInRow: false, IsLastInRow: true},
	}
}

func (n Note) Print() {
	ui.DrawHorizontalLine()
	fmt.Println("\t Id:          " + strconv.Itoa(n.id))
	fmt.Println("\t Title:       " + n.title)
	fmt.Println("\t Description: " + n.description)
	fmt.Println("\t CreatedAt:   " + n.createdAt.String())
	fmt.Println("\t UpdatedAt:   " + n.updatedAt.String())
	ui.DrawHorizontalLine()
}
