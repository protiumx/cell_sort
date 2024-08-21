package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Cell struct {
	current_pos    int
	target_pos     int
	left_neighbor  *Cell
	right_neighbor *Cell
	is_moving      bool
	value          int
}

func (c *Cell) move_to_right() {
	if c.right_neighbor == nil {
		return
	}

	r := c.right_neighbor
	c.right_neighbor = r.right_neighbor
	if r.right_neighbor != nil {
		r.right_neighbor.left_neighbor = c
	}

	r.left_neighbor = c.left_neighbor
	c.left_neighbor.right_neighbor = r

	r.right_neighbor = c
	c.left_neighbor = r
}

func (c *Cell) Value() int {
	return c.value
}

func new_cell(value int, current_pos int) *Cell {
	return &Cell{
		value:       value,
		current_pos: current_pos,
	}
}

func (c *Cell) should_move_right() bool {
	if c.right_neighbor == nil {
		return false
	}

	return c.value > c.right_neighbor.value
}

func (c *Cell) String() string {
	var out strings.Builder
	out.WriteString("Cell{value: ")
	out.WriteString(strconv.Itoa(c.value))
	out.WriteString(", pos: ")
	out.WriteString(strconv.Itoa(c.current_pos))
	if c.left_neighbor != nil {
		out.WriteString(", left: ")
		out.WriteString(strconv.Itoa(c.left_neighbor.value))
	}

	if c.right_neighbor != nil {
		out.WriteString(", right: ")
		out.WriteString(strconv.Itoa(c.right_neighbor.value))
	}

	out.WriteString("}")
	return out.String()
}

func build_list(l []int) ([]*Cell, *Cell) {
	ret := make([]*Cell, len(l))
	current := new_cell(l[0], 0)
	start := &Cell{right_neighbor: current}
	current.left_neighbor = start

	ret[0] = current
	for i := 1; i < len(l); i++ {
		cell := new_cell(l[i], i)
		ret[i] = cell

		cell.left_neighbor = current
		current.right_neighbor = cell
		current = cell
	}

	return ret, start
}

func print(start *Cell) {
	for start != nil {
		fmt.Println(start)
		start = start.right_neighbor
	}
}

func sort_cells(start *Cell, cells []*Cell) {
	should_sort := true
	step := 1
	for should_sort {
		should_sort = false
		for _, cell := range cells {
			if !cell.should_move_right() {
				continue
			}

			should_sort = true
			cell.move_to_right()
			step++
			break
		}
	}
}

func main() {
	cells, start := build_list([]int{10, 38, 1, 3, 2, 0, 99})
	fmt.Println("before")
	print(start.right_neighbor)

	fmt.Println("\nafter")
	sort_cells(start, cells)
	print(start.right_neighbor)
}
