package cligobrr

import "fmt"
import "strings"

type TableFields struct {
	Cols uint8
	Pad  uint8
}

type Table struct {
	TableFields
	lens []uint8
	rows [][]string
}

func tableNew(fields TableFields) (*Table, error) {
	if fields.Cols == 0 {
		return nil, errTableColsRequired()
	}

	if fields.Pad == 0 {
		fields.Pad = tablePadDefault
	}

	table := Table{
		TableFields: fields,
		lens:        make([]uint8, fields.Cols, fields.Cols),
	}

	return &table, nil
}

func (self *Table) Add(row []string) error {
	rowLen := uint8(len(row))
	if rowLen != self.Cols {
		return errTableRowIncorrectCols(self.Cols)
	}

	for i, val := range row {
		valLen := uint8(len(val))
		if valLen > self.lens[i] {
			self.lens[i] = valLen
		}
	}

	self.rows = append(self.rows, row)

	return nil
}

func (self *Table) normalize() {
	for _, row := range self.rows {
		for i, cell := range row {
			cellLen := uint8(len(cell))
			padLen := self.lens[i] - cellLen
			if padLen > 0 {
				pad := strings.Repeat(" ", int(padLen))
				row[i] = fmt.Sprintf("%s%s", cell, pad)
			}
		}
	}
}

func (self *Table) ToString() string {
	var output []string

	self.normalize()

	padding := strings.Repeat(" ", int(self.Pad))

	for _, row := range self.rows {
		output = append(output, strings.Join(row, padding))
	}

	return strings.Join(output, "\n")
}
