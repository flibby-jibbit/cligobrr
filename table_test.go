package cligobrr

import "testing"
import "github.com/stretchr/testify/assert"

func TestTableNew(t *testing.T) {
	assert := assert.New(t)

	fields := TableFields{
		Cols: 3,
		Pad:  3,
	}

	table, err := tableNew(fields)
	assert.NotNil(table)
	assert.Nil(err)
	assert.Equal(fields.Cols, table.Cols)
	assert.Equal(fields.Pad, table.Pad)
	assert.Equal(int(fields.Cols), len(table.lens))
}

func TestTableNewZeroCols(t *testing.T) {
	assert := assert.New(t)

	fields := TableFields{}

	table, err := tableNew(fields)
	assert.Nil(table)
	assert.NotNil(err)
}

func TestTableNewZeroPad(t *testing.T) {
	assert := assert.New(t)

	fields := TableFields{
		Cols: 3,
	}

	table, err := tableNew(fields)
	assert.NotNil(table)
	assert.Nil(err)
	assert.Equal(tablePadDefault, table.Pad)
	assert.Equal(0, len(table.rows))
}

func TestTableAdd(t *testing.T) {
	assert := assert.New(t)

	fields := TableFields{
		Cols: 3,
	}

	table, err := tableNew(fields)
	assert.Nil(err)

	row := []string{"one", "two", "three"}
	err = table.Add(row)
	assert.Nil(err)
	assert.Equal(1, len(table.rows))
	assert.Equal(row, table.rows[0])
}

func TestTableAddIncorrectCols(t *testing.T) {
	assert := assert.New(t)

	fields := TableFields{
		Cols: 3,
	}

	table, err := tableNew(fields)
	assert.Nil(err)

	row := []string{"one", "two"}
	err = table.Add(row)
	assert.NotNil(err)
	assert.Equal(0, len(table.rows))
}

func TestTableNormalize(t *testing.T) {
	assert := assert.New(t)

	fields := TableFields{
		Cols: 3,
	}

	table, err := tableNew(fields)
	assert.Nil(err)

	row1 := []string{"one", "two", "three"}
	row2 := []string{"four", "five", "six"}
	err = table.Add(row1)
	assert.Nil(err)
	err = table.Add(row2)
	assert.Nil(err)
	assert.Equal(2, len(table.rows))

	table.normalize()
	assert.Equal([]string{"one ", "two ", "three"}, table.rows[0])
	assert.Equal([]string{"four", "five", "six  "}, table.rows[1])
}

func TestTableToString(t *testing.T) {
	assert := assert.New(t)

	fields := TableFields{
		Cols: 3,
	}

	table, err := tableNew(fields)
	assert.Nil(err)

	row1 := []string{"one", "two", "three"}
	row2 := []string{"four", "five", "six"}
	err = table.Add(row1)
	assert.Nil(err)
	err = table.Add(row2)
	assert.Nil(err)
	assert.Equal(2, len(table.rows))

	expected := "one      two      three\nfour     five     six  "
	actual := table.ToString()
	assert.Equal(expected, actual)
}
