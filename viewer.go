package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

const (
	SortAsc int = iota
	SortDesc
)

// Sort list by converting compared values into Integer
func sortByIntColumn(list [][]string, sortIndex, sortDirection int) [][]string {
	sort.Slice(list[:], func(i, j int) bool {
		iInt, _ := strconv.Atoi(list[i][sortIndex])
		jInt, _ := strconv.Atoi(list[j][sortIndex])
		if sortDirection == SortDesc {
			return iInt > jInt
		} else {
			return iInt < jInt
		}
	})
	return list
}

// Sort list by converting compared values into Float64
func sortByFloatColumn(list [][]string, sortIndex, sortDirection int) [][]string {
	sort.Slice(list[:], func(i, j int) bool {
		iFloat, _ := strconv.ParseFloat(list[i][sortIndex], 64)
		jFloat, _ := strconv.ParseFloat(list[j][sortIndex], 64)
		if sortDirection == SortDesc {
			return iFloat > jFloat
		} else {
			return iFloat < jFloat
		}
	})
	return list
}

// Load table data from CSV file.
func loadTableData(id string, sortIndex, sortDirection int) [][]string {
	rows := ReadCSV(fmt.Sprintf("rqmetric_output_%v.csv", id))
	var reqs [][]string

	for index, row := range rows {
		// remove the original csv header
		if index == 0 {
			continue
		}
		reqs = append(reqs, row)
	}

	// apply sorting based on column index (descending)
	switch sortIndex {
	case 0: // return unsorted
		reqs = reqs
	case 3: // sort by Avg. time
		reqs = sortByFloatColumn(reqs, sortIndex, sortDirection)
	default:
		reqs = sortByIntColumn(reqs, sortIndex, sortDirection)
	}

	// add the header
	var header [][]string
	header = append(header, []string{"URL", "Min. (ms)", "Max. (ms)", "Avg. (ms)", "Count", "2XX", "3XX", "4XX", "5XX"})
	return append(header, reqs...)
}

func StartViewer(id string) {
	requests := loadTableData(id, 4, SortDesc) // sort by Count by default

	viewer := tview.NewApplication()
	table := tview.NewTable()
	cols, rows := len(RequestCsvHeader()), len(requests)

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			color := tcell.ColorWhite
			switch {
			case r == 0:
				color = tcell.ColorYellow
			case c == 0 && r > 0:
				color = tcell.ColorGreen
			}
			cell := tview.NewTableCell(requests[r][c]).SetTextColor(color).SetAlign(tview.AlignLeft).SetExpansion(1)
			table.SetCell(r, c, cell)
		}
	}

	table.SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			viewer.Stop()
		}
	}).SetSelectedFunc(func(row, col int) {
		log.Println(row, col)
	})

	if err := viewer.SetRoot(table, true).SetFocus(table).Run(); err != nil {
		log.Panic(err)
	}
}
