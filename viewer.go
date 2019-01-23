package main

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// TableData will holds the table data and its header.
type TableData struct {
	data   [][]string
	header []string
}

// LoadData data from specified CSV file and set the TableData.data.
func (td *TableData) LoadData(fileName string) {
	rows := ReadCSV(fileName)

	for index, row := range rows {
		// remove the original csv header
		if index == 0 {
			continue
		}
		td.data = append(td.data, row)
	}
}

// SetHeader sets the TableData.header.
func (td *TableData) SetHeader(header []string) {
	td.header = header
}

// SortData do sorting to the TableData.data.
// If the sortIndex is 0 which is the column of URL, ignore it.
func (td *TableData) SortData(sortIndex, sortDirection int) {
	// don't sort if sort index is the URL column.
	if sortIndex == 0 {
		return
	}

	switch sortIndex {
	case 3: // sort by Avg. time
		td.data = SortByFloatColumn(td.data, sortIndex, sortDirection)
	default:
		td.data = SortByIntColumn(td.data, sortIndex, sortDirection)
	}
}

// DataForTable returns complete data with its header.
func (td *TableData) DataForTable() [][]string {
	var tableRows [][]string
	tableRows = append(tableRows, td.header)
	return append(tableRows, td.data...)
}

// fillTable fills the table cells.
func fillTable(table *tview.Table, cols, rows int, data [][]string) {
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			color := tcell.ColorWhite
			switch {
			case r == 0:
				color = tcell.ColorYellow
			case c == 0 && r > 0:
				color = tcell.ColorGreen
			}
			cell := tview.NewTableCell(data[r][c]).
				SetTextColor(color).
				SetAlign(tview.AlignLeft).
				SetExpansion(1).
				SetMaxWidth(800)
			table.SetCell(r, c, cell)
		}
	}
}

// StartViewer start the data viewer application.
func StartViewer(id string) {
	tableData := &TableData{}
	tableData.SetHeader([]string{"URL", "Min. (ms)", "Max. (ms)", "Avg. (ms)", "Count", "2XX", "3XX", "4XX", "5XX"})
	tableData.LoadData(fmt.Sprintf("rqmetric_output_%v.csv", id))
	tableData.SortData(4, SortDesc) // by default we sort it by request count descending.

	data := tableData.DataForTable()

	viewer := tview.NewApplication()
	table := tview.NewTable()
	cols, rows := len(tableData.header), len(data)

	fillTable(table, cols, rows, data)

	table.SetFixed(1, 1).
		SetSelectable(false, true).
		Select(0, 4). // select request count column by default.
		SetSelectedStyle(tcell.ColorDefault, tcell.ColorDarkSlateGray, 0).
		SetDoneFunc(func(key tcell.Key) {
			if key == tcell.KeyEscape {
				viewer.Stop()
			}
		}).
		SetSelectedFunc(func(row, col int) {
			// when Enter key pressed on selected column, sort table data based on that column.
			viewer.QueueUpdateDraw(func() {
				tableData.SortData(col, SortDesc)
				fillTable(table, cols, rows, tableData.DataForTable())
				table.ScrollToBeginning()
			})
		})

	if err := viewer.SetRoot(table, true).SetFocus(table).Run(); err != nil {
		log.Panic(err)
	}
}
