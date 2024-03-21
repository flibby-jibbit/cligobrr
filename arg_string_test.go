package cligobrr

import "testing"
import "github.com/stretchr/testify/assert"

func TestStringArgNew(t *testing.T) {
	assert := assert.New(t)

	fields := ArgFields{
		Name: "sarg",
	}

	arg, err := StringArgNew(fields)
	assert.Nil(err)
	assert.Equal(kindString, arg.GetKind())
}
