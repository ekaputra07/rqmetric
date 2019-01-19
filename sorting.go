package main

import (
	"sort"
	"strconv"
)

// Constants of sorting direction
const (
	SortDesc int = iota
)

// SortByIntColumn sort list by converting compared values into Integer
func SortByIntColumn(list [][]string, sortIndex, sortDirection int) [][]string {
	sort.Slice(list[:], func(i, j int) bool {
		iInt, _ := strconv.Atoi(list[i][sortIndex])
		jInt, _ := strconv.Atoi(list[j][sortIndex])
		if sortDirection == SortDesc {
			return iInt > jInt
		}
		return iInt < jInt
	})
	return list
}

// SortByFloatColumn sort list by converting compared values into Float64
func SortByFloatColumn(list [][]string, sortIndex, sortDirection int) [][]string {
	sort.Slice(list[:], func(i, j int) bool {
		iFloat, _ := strconv.ParseFloat(list[i][sortIndex], 64)
		jFloat, _ := strconv.ParseFloat(list[j][sortIndex], 64)
		if sortDirection == SortDesc {
			return iFloat > jFloat
		}
		return iFloat < jFloat
	})
	return list
}
